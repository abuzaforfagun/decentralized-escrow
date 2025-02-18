# Decentralized Escrow

## Flow:

- Create Product: Seller submit product creation request from the frontend to the smart contract. Smart contract return the product id. Frontend call the backend API using the product id and transaction hash. Backend check the transaction status and verify the product id. When everything looks good, it add the product to the database.
- Update Price: Seller update the price from the frontend, frontend trigger a contract call. Smart contract verify the seller, and update the price, also trigger an event. Backend listen to that event and update the price in the database.
- Update stock: Seller update the stock of the product from the frontend, it triggers a contract call, contract update the stock and fire an event. Backend listen to that event and update the stock.
- Create order: Buyer interact with the chian from the frontend, pay the fee using their wallet, on sucessful transaction, smart contract trigger an event, backend listen to the event and process. During order creation, contract keep the money until recive confirmation.
- Order recival confirmation: When user receive the order, he confirm the recival, and money get paid to the seller.
- No recival in the due time: When product is not recived within X days, backend trigger a contract call, and change the status to Disputed.
- Resolve dispute: In the X amount of days, seller need to prove the product delivery or the customer need to confirm recival, otherwise backend need to trigger another contract call to refund the money to the buyer. On recival confirmation another contract call change the status to Resolved and send the money to the buyer.
