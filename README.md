# Go Concurrent Email Campaigner



A high-performance, concurrent email dispatching utility written in Go. This tool reads recipient information from a CSV file, populates a custom email template, and sends emails through an SMTP server using a pool of worker goroutines.

# Features

 
Producer-Consumer Architecture: Decouples file reading from email sending for maximum efficiency.

Concurrent Workers: Uses a configurable pool of workers to handle multiple emails simultaneously.

Dynamic Templating: Utilizes Go's html/template for professional, personalized email bodies.

Wait Group Synchronization: Ensures the application only exits after every recipient has been processed.

# Project Structure


main.go: Entry point that initializes channels, workers, and the producer.

producer.go: Contains logic to parse email.csv and feed the recipient channel.

consumer.go: Implements the worker logic for SMTP communication.

email.tmpl: The HTML/Text template for the email content.

email.csv: Data source containing recipient names and email addresses.
