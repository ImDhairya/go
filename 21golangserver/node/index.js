import express from "express";

const PORT = 4000;

const app = express();

app.use(express.json());
app.use(express.urlencoded({extended: true}));

app.get("/", (req, res) => {
  res.status(200).send("Welcome to LearnCodeonline server");
});

app.get("/get", (req, res) => {
  res.status(200).json({message: "Hello from learnCodeonline.in"});
});

app.post("/post", (req, res) => {
  let myJson = req.body; // your JSON

  res.status(200).send(myJson);
});

app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});

app.listen(PORT, () => {
  console.log("The app is listening to port ", PORT);
});
