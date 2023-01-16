Tema 16 - Stress Test Chaos

* Requisitos:
	- Docker
	- Docker-compose
	- Maven
	- Scala
	- Jenkins
	
* Configurando o Jenkins:
		
		1.	Criar novo item
		2.	Colocar o nome do Job e escolher a opção 'Pipeline'
		3.	Em seguida no campo Pipeline definition escolher a opção 'Pipeline script from SCM'
		4.	No campo SCM escolher a opção 'Git'
		5.	No campo Repository URL inserir: 'https://github.com/julianooen/DevOpsTema16.git'
		6.	Em Branch Specifier preencher com '*/master'
		7.	Em Script Path preencher com 'jenkins/Jenkinsfile'
	
* Resultados:
	
	- Após rodar o job, é possivel acessar e visualizar os resultados dos testes clicando na opção 'gatling'
	- Também é possivel acessar os resultado indo em var/lib/jenkins/workspace e procurar pela pasta do projeto, 
	dentro dela vai conter o index.html para visualizar os resultados.

* Acessando o serviço:

	Operações:
	- Soma: http://localhost:8080/calc/sum/10/5
	- Subtração: http://localhost:8080/calc/sub/10/5
	- Divisão: http://localhost:8080/calc/div/10/5
	- Multiplicação: http://localhost:8080/calc/mul/10/5
	- Histórico: http://localhost:8080/calc/history
