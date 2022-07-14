from flask import Flask, jsonify
import sqlite3

app = Flask(__name__)

@app.route("/", methods=["GET"])
def home():
  try:
    dbconnection = sqlite3.connect("../test.db")
    cursor = dbconnection.cursor()
    cursor.execute("SELECT * FROM test")
    data = cursor.fetchall()

    data2 = []
    for row in data:
      data2.append({"id": row[0], "name": row[1]})

    result = []
    for x in range(10):
      result.extend(data2)

    return jsonify(result)

  except sqlite3.Error as error:
    print('Error occured - ', error)
    
  finally:
    dbconnection.close()

app.run(debug=False)