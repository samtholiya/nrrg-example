import "reflect-metadata";
import { createConnection, Connection } from "typeorm";
import * as express from 'express';
import * as bodyParser from 'body-parser'
import { Md5 } from "md5-typescript";
import { Api } from "./api";
import { IServicePanel } from "./namespace/initializer";
import { DatabaseService } from "./service/database.service";


console.log(Md5.init('root'));
const app: express.Application = express();
const port: number = Number(process.env.PORT) || 3000;


app.use(bodyParser.json());
app.use((req, res, next) => {
    console.log(req.url);
    next();
});

//app.use('/register', RegisterController);
app.use('/api', Api);

console.log("Connecting to Database");
DatabaseService.connectDatabase()
    .then((connection: Connection) => {
        
        console.info("Database connected");
        console.info('Initializing services');
        var servicePanels = IServicePanel.GetImplementations();
        for (var x = 0; x < servicePanels.length; x++) {
            console.info('Initializing ' + servicePanels[x].name);
            let panel = new servicePanels[x]();
            panel.init();
        }
        console.info('Initialization complete');

        app.listen(port, '0.0.0.0', () => {
            console.log(`Listening at http://localhost:${port}/`);
        });

    }).catch(reason=>{
        console.log(reason);
    })
