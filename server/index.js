const express = require("express");
const bodyParser = require("body-parser");

const App = express();

function Logger(ip, body) {
	console.log(`FROM ${ip}`);
	console.log(body);
}

App.use(express.json({ limit: "10mb"}));

App.post("/", (req,res) => {
	Logger(req.id, req.body);
	res.status(200).json({ status: "success" });
});

App.listen(8080, () => {
	console.log("Listening on " + 8080);
});
