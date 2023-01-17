import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._
import com.typesafe.config._

class TestSimulation extends Simulation {
  
  val untrackedHeader = "1"
  val props = ConfigFactory.load()
  val envPort = props.getString("portApp")

  val httpConf = http
  .baseUrl("http://localhost:" + envPort) 
  .acceptHeader("text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
  .doNotTrackHeader(untrackedHeader)
  .acceptLanguageHeader("en-US,en;q=0.5")
  .acceptEncodingHeader("gzip, deflate")
  .disableFollowRedirect
  .userAgentHeader("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0")

  val scn = scenario("Test")
          .exec(http("History") 
            .get("/calc/history")
            .check(status.is(200))
          )
    
  setUp(
      scn.inject(constantUsersPerSec(100).during(20 seconds))
  ).protocols(httpConf)

}