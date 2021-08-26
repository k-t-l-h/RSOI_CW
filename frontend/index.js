let express = require('express');
let cors = require('cors');

let app = express();
const port = 3000;

app.get('/*', (req, res) => {
    res.sendFile('./main/main.html');
})

app.listen(port, () => {
    console.log(`App listening at http://localhost:${port}`)
})
