
# Kafka Producer and Consumer Service (Go)

Este projeto demonstra uma arquitetura baseada em eventos (EDA - Event Driven Architecture) com serviços desenvolvidos em Go utilizando Kafka como sistema de mensageria.

## Visão Geral

A aplicação é composta por dois microsserviços principais:

- **Producer Service**: recebe requisições HTTP e publica eventos no tópico `transactions` do Kafka.
- **Consumer Service**: consome os eventos de forma assíncrona e os distribui internamente para múltiplos handlers por meio de um `Dispatcher` in-memory.

## Arquitetura

![Captura de tela de 2025-04-30 17-38-36](https://github.com/user-attachments/assets/11db1de3-50a4-4c47-a28d-8e65963561a7)


- O `Producer Service` publica eventos no Kafka.
- O `Consumer Service` escuta esses eventos e, ao receber uma mensagem, despacha para diferentes `handlers` registrados em memória.
- Os `handlers` representam unidades de negócio, como:
  - `Process`: processamento da transação.
  - `FraudDetection`: simulação de detecção de fraude.


