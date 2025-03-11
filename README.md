# backend-challenge-from-devops-perspective

🚀 Como Funciona
1️⃣ Build e Push da Imagem Docker
A pipeline GitHub Actions é acionada quando há um commit na branch develop. O workflow realiza os seguintes passos:

*Faz build da imagem da aplicação Java.
*Faz push da imagem para o Docker Hub.
*Envia um log de sucesso.
