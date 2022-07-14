import http from "node:http";
import Database from "better-sqlite3";

const hostname = "127.0.0.1";
const port = 3000;

const db = new Database("../test.db");
const stm = db.prepare("SELECT * from test");

const server = http.createServer((req, res) => {
  try {
    let data = [];
    for (let i = 0; i < 10; i++) {
      data.push(...stm.all());
    }

    res.setHeader("Content-Type", "application/json");
    res.writeHead(200);
    res.end(JSON.stringify(data));
  } catch (e) {
    res.statusCode = 500;
    res.end(JSON.stringify(e));
  }
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
