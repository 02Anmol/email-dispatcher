# Go Email Dispatcher 📧

A concurrent email processing system built with Go. This project reads recipient data from a CSV file and uses a worker pool pattern to process and "send" emails efficiently.

## 🚀 Features
* **Concurrent Processing**: Uses Go routines to read and process data simultaneously.
* **Worker Pool**: Scalable design using multiple workers to handle high-volume lists.
* **Safe Communication**: Implements Go channels for memory-safe data transfer between the producer and consumers.
* **CSV Integration**: Automatically parses recipient names and emails from external files.

## 🏗️ Architecture
The project follows a **Producer-Consumer** pattern:
1. **Producer (`producer.go`)**: Reads the `sample.csv` file, converts rows into `Recipient` structs, and sends them into a channel.
2. **Channel**: Acts as a conveyor belt to hold data for the workers.
3. **Consumers (`consumer.go`)**: A pool of workers that listen to the channel and process the emails.
4. **WaitGroup**: Ensures the main program waits for all workers to finish before exiting.

## 🛠️ Prerequisites
* **Go**: Version 1.23 or higher.
* **Mailpit**: (Optional) For local SMTP testing.

## 📋 Project Structure
* `main.go`: Entry point that initializes the channel and orchestrates the workers.
* `producer.go`: Contains the `loadRecipient` logic for file parsing.
* `consumer.go`: Contains the `emailWorker` logic for handling tasks.
* `sample.csv`: Data source containing name and email information.

## ⚙️ How to Run
1. Ensure your `sample.csv` is in the root directory.
2. Run the entire package using:
   ```bash
   go run .