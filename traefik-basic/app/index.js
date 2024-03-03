const app = require("express")();
const containerNumber = process.env.CONTAINER_NUMBER || 'no number';

app.get("/", (req, res) => {
    res.send(`Hello from a lightweight container ${containerNumber}`);
});

app.get('/error-403', (req, res) => {
    res.status(403).send('403 Forbidden: Access is denied.');
});

app.get('/error-404', (req, res) => {
    res.status(404).send('404 Not Found: The resource you requested could not be found.');
});

app.get('/error-429', (req, res) => {
    res.status(429).send('429 Too Many Requests: You have sent too many requests in a given amount of time.');
});

app.get('/error-500', (req, res) => {
    res.status(500).send('500 Internal Server Error: The server encountered an unexpected condition.');
});

app.get('/error-501', (req, res) => {
    res.status(501).send('501 Not Implemented: The server does not support the functionality required to fulfill the request.');
});

app.get('/error-502', (req, res) => {
    res.status(502).send('502 Bad Gateway: The server received an invalid response from the upstream server.');
});

app.get('/error-503', (req, res) => {
    res.status(503).send('503 Service Unavailable: The server is currently unable to handle the request due to temporary overloading or maintenance.');
});

// Catch-all for any other GET requests not handled above
app.get('*', (req, res) => {
    res.status(404).send('404 Not Found: The resource you requested could not be found.');
});

app.listen(9999, () => console.log("Listening on port 9999"));
