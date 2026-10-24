# CEP Finder v1.0

> *"Busca rápida, precisa e eficiente de dados de CEP com arquitetura multi-linguagem."*

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Java](https://img.shields.io/badge/Java-ED8B00?style=for-the-badge&logo=java-applet&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![Status](https://img.shields.io/badge/Status-Conclu%C3%ADdo-success?style=for-the-badge)

---

## Sobre o Projeto

O **CEP Finder 1.0** é uma aplicação web desenvolvida para consulta e validação de Códigos de Endereçamento Postal (CEP). O projeto adota uma arquitetura distribuída e eficiente, aproveitando os pontos fortes de cada linguagem:

* **Go (Golang):** Responsável pela alta performance no processamento e intermediação de requisições / serviços de consulta.
* **Java:** Utilizado no backend/serviços para gerenciamento da lógica de negócios e estruturação de dados.
* **JavaScript:** Responsável pela interatividade na interface gráfica, manipulação dinâmica do DOM e aplicação de máscaras de entrada em tempo real.

---

## Divisão do Desenvolvimento

* **Serviço em Go & Máscara de Entrada (JS):** Desenvolvido por **[@JV74](https://github.com/joao7834)**
* **Serviço em Java & Estrutura Base:** Desenvolvido por **[@FXVIZFF2](https://github.com/FXVIZFF2)**

---

## Funcionalidades

* **Consulta Instantânea:** Localização rápida de endereços a partir do CEP informado.
* **Mascara Automática (JS):** Formatação dinâmica do campo de texto no formato `00000-000` enquanto o usuário digita.
* **Performance e Baixa Latência:** Backend otimizado em Go para consultas eficientes.
* **Interface Amigável:** Design limpo para facilidade de navegação e uso.

---

## Tecnologias Utilizadas

* **Golang** — Gateway/Serviço de alta concorrência e performance.
* **Java** — Regras de negócio e integração de serviços.
* **JavaScript (ES6+)** — Manipulação dinâmica do DOM e máscaras de entrada.
* **HTML5 & CSS3** — Estrutura e estilização visual da aplicação.

---

## Como Executar o Projeto

### Pré-requisitos
* **Go** (v1.26+)
* **Java JDK** (17 ou superior)
* Um navegador web moderno

### Passo a Passo

1. Clone o repositório:
   ```bash
   git clone [https://github.com/joao7834/cep-finder-go-java.git](https://github.com/joao7834/cep-finder-go-java.git)
   cd cep-finder-go-java
