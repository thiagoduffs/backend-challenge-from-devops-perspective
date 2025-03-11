# backend-challenge-from-devops-perspective


# 📂 Estrutura do Projeto

# **SRE Challenge - Java Application on AWS EKS**

Este repositório contém uma aplicação **Java Spring Boot** configurada para rodar em um ambiente **AWS EKS** com deploy automatizado via **GitHub Actions** e monitoramento com **AWS CloudWatch**. O objetivo desse projeto é demonstrar **boas práticas de SRE** em **containerização, deploy automatizado e observabilidade**.

---

## **🛠️ Tecnologias Utilizadas**
A stack do projeto foi construída com as seguintes tecnologias:

### **📌 Backend**
- Java 17 com **Spring Boot**
- API REST simples (Hello World)

### **📌 Infraestrutura**
- **Docker** para containerizar a aplicação
- **Helm** para gerenciar os manifestos do Kubernetes
- **AWS EKS** como orquestrador de containers
- **AWS CloudWatch** para monitoramento do cluster
- **AWS SNS** para notificações de eventos críticos

### **📌 CI/CD**
- **GitHub Actions** para build e push da imagem Docker
- **Helm Chart** para deploy no Kubernetes

---

## 📂 Estrutura do Projeto

```bash
.
├── app                          # Código da aplicação Java
│   ├── mvnw                     # Wrapper do Maven
│   ├── mvnw.cmd                 # Wrapper do Maven (Windows)
│   ├── pom.xml                  # Arquivo de configuração do Maven
│   ├── src                      # Código-fonte da aplicação
│   │   ├── main
│   │   │   ├── java/com/example/app
│   │   │   │   ├── AppAController.java
│   │   │   │   ├── AppApplication.java
│   │   │   └── resources
│   │   │       ├── application.properties  # Configuração da aplicação
│   │   └── test/java/com/example/app
│   │       ├── AppApplicationTests.java    # Testes da aplicação
│   ├── target                              # Artefatos gerados pelo Maven
│       ├── app-0.0.1-SNAPSHOT.jar
│       ├── app-0.0.1-SNAPSHOT.jar.original
│       ├── classes
│       │   ├── application.properties
│       │   └── com/example/app
│       │       ├── AppAController.class
│       │       ├── AppApplication.class
│       ├── generated-sources/annotations
│       ├── generated-test-sources/test-annotations
│       ├── maven-archiver/pom.properties
│       ├── maven-status/maven-compiler-plugin
│       │   ├── compile/default-compile
│       │   │   ├── createdFiles.lst
│       │   │   ├── inputFiles.lst
│       │   ├── testCompile/default-testCompile
│       │       ├── createdFiles.lst
│       │       ├── inputFiles.lst
│       └── test-classes/com/example/app
│           ├── AppApplicationTests.class
├── Dockerfile                        # Configuração da imagem Docker
├── README.md                         # Documentação do projeto
├── sre-challenge-java                 # Helm Chart para deploy no Kubernetes
│   ├── Chart.yaml                     # Definições do Helm Chart
│   ├── values.yaml                     # Configurações do Helm Chart
│   ├── templates                      # Manifestos Kubernetes
│   │   ├── deployment.yaml             # Configuração do Deployment no Kubernetes
│   │   ├── service.yaml                # Configuração do Service no Kubernetes
│   │   ├── ingress.yaml                # Configuração do Ingress (opcional)
│   │   ├── namespace.yaml              # Namespace da aplicação
│   │   ├── hpa.yaml                    # Configuração do Autoscaler
│   │   ├── serviceaccount.yaml         # Configuração da Service Account
│   │   ├── _helpers.tpl                # Helpers do Helm
│   │   ├── tests/test-connection.yaml  # Testes do Helm
└── terraform-eks                      # Infraestrutura como código (IaC) para criar o cluster EKS
    ├── eks.tf                          # Configuração do EKS
    ├── provider.tf                      # Configuração do Provider AWS
    ├── security_groups.tf               # Regras de firewall do cluster
    ├── nodegroup.tf                     # Configuração do grupo de nós
    ├── network.tf                       # Configuração de rede
    ├── elb.tf                           # Configuração do Load Balancer
    ├── iam.tf                           # Permissões IAM
    ├── outputs.tf                        # Saídas do Terraform
    ├── variables.tf                      # Variáveis do Terraform




