# GST Billing System

A simple and efficient billing system designed to handle product purchases, generate bills, and maintain a history of transactions with GST calculations.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Features

- Generate bills for multiple products with quantity support.
- Calculate total, subtotal, and GST automatically.
- Save and retrieve billing history for users.
- RESTful API for easy integration with other applications.

## Technologies

- **Go**: Programming language for the backend.
- **Gin**: Web framework for building APIs.
- **GORM**: ORM library for interacting with PostgreSQL.
- **PostgreSQL**: Database for storing product and billing data.

## Installation

### Prerequisites

- Go installed on your machine.
- PostgreSQL database set up and running.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gst-billing-system.git
   cd gst-billing-system
