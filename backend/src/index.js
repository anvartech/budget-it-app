import bugsnag from 'bugsnag';
import express from 'express';
import bodyParser from 'body-parser';
import logger from 'morgan';

const app = express();
const environment = process.env.NODE_ENV;

if (environment !== 'production') {
    require('dotenv').config();
    app.use(logger('dev'));
}
bugsnag.register(process.env.BUGSNAG_API)

app.use(bugsnag.requestHandler);
app.use(bugsnag.errorHandler);

const SECRET_KEY = process.env.SECRET_KEY;
if (SECRET_KEY) {
    app.set('JWTSecret', SECRET_KEY);
} else {
    throw new Error('No Secret Key');
}

app.use(bodyParser.urlencoded({
    extended: false,
}));
app.use(bodyParser.json({
    type: 'application/json',
}));









const port = parseInt(process.env.PORT, 10) || 3000;
app.set('port', port);

const server = http.createServer(app);
server.listen(port);

export default app;
