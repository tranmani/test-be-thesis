from flask import Flask, jsonify
import sqlite3

app = Flask(__name__)

@app.route("/", methods=["GET"])
def home():
  dbconnection = sqlite3.connect("../test.db")
  cursor = dbconnection.cursor()
  cursor.execute("SELECT * FROM test")
  data = cursor.fetchall()

  result = []
  for x in range(10):
    result.extend(data)

  return jsonify(result)

app.run(debug=True)