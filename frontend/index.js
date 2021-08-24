let express = require('express');
let cors = require('cors');

let app = express();
const port = 3000;


app.use(cors());

app.listen(port, () => {
    console.log(`Fish app listening at http://localhost:${port}`)
})
