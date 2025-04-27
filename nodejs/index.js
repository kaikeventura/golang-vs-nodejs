const express = require('express');
const app = express();
const port = 8080;

app.use(express.json());

function fibonacci(n) {
  if (n <= 1) {
    return n;
  }
  return fibonacci(n - 1) + fibonacci(n - 2);
}

app.post('/fibonacci', (req, res) => {
  const { number } = req.body;

  if (typeof number !== 'number') {
    return res.status(400).json({ error: "Campo 'number' precisa ser um número." });
  }

  const results = [];
  for (let i = 0; i < 1000; i++) {
    results.push(fibonacci(number));
  }

  return res.status(200).json(results);
});

app.listen(port, () => {
  console.log(`Servidor rodando na porta 8080`);
});

