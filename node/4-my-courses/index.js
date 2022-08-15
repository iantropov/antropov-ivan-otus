const express = require('express');
const app = express();

app.set('view engine', 'pug');

app.use((req, res, next) => {
    console.log(`[${new Date().toISOString()}] - Request received`);
    next();
})

app.get('/', (req, res) => {
    res.render('index', { title: 'Hello title', message: 'Hello message' });
});

app.get('/test', (req, res) => {
    res.send("Hello world! You've sent me - " + req.query.name);
});

app.listen(3000, () => console.log('Hello from 3000'));
