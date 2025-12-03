<div align="center">

# 📧 Spam Email Classifier

### Intelligent Email Classification with Deep Learning

[![Python](https://img.shields.io/badge/Python-3.10-blue.svg)](https://www.python.org/)
[![Go](https://img.shields.io/badge/Go-1.21-00ADD8.svg)](https://golang.org/)
[![TensorFlow](https://img.shields.io/badge/TensorFlow-2.14-FF6F00.svg)](https://www.tensorflow.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED.svg)](https://www.docker.com/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A production-ready web application that uses advanced bidirectional LSTM neural networks to classify emails as spam or legitimate with **97% accuracy**.

[Features](#-features) • [Quick Start](#-quick-start-with-docker) • [API](#-api-endpoints) • [Model](#-model-architecture) • [Contributing](#-contributing)

---

</div>

## ✨ Features

- 🚀 **Real-time Classification** - Instant email spam detection
- 🎯 **High Accuracy** - 97% accuracy on test dataset
- 🌐 **RESTful API** - Easy integration with any frontend
- 🐳 **Docker Support** - One-command deployment
- 📱 **Responsive UI** - Beautiful, mobile-friendly interface
- ⚡ **Fast Inference** - Go backend with Python ML engine
- 🔒 **Production Ready** - Built with best practices

## 🏗️ Architecture

<div align="center">

```
┌─────────────┐      ┌──────────────┐      ┌─────────────────┐
│   Frontend  │ ───▶ │  Go Backend  │ ───▶ │ Python ML Engine│
│  (HTML/JS)  │ ◀─── │   (REST API) │ ◀─── │  (TensorFlow)   │
└─────────────┘      └──────────────┘      └─────────────────┘
```

</div>

| Component | Technology | Purpose |
|-----------|-----------|---------|
| **Frontend** | HTML5, CSS3, JavaScript | User interface |
| **Backend** | Go (Golang) | High-performance API server |
| **ML Engine** | Python + TensorFlow/Keras | Model inference |
| **Model** | Bidirectional LSTM | Deep learning classifier |
| **Deployment** | Docker + Docker Compose | Containerization |

## 📋 Prerequisites

### Option 1: Docker (Recommended)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) 20.10+
- [Docker Compose](https://docs.docker.com/compose/) 2.0+

### Option 2: Manual Setup
- [Go](https://golang.org/dl/) 1.21+
- [Python](https://www.python.org/downloads/) 3.10+
- Required Python packages (see `model/requirements.txt`)

## 🚀 Quick Start with Docker

```bash
# 1. Clone the repository
git clone https://github.com/SahithiKokkula/Spam-Email-Classifier.git
cd Spam-Email-Classifier

# 2. Build and run with Docker Compose
docker-compose up --build

# 3. Open your browser
# Navigate to: http://localhost:8080
```

**That's it!** 🎉 Your spam classifier is now running!

## Manual Setup

### Backend Setup

1. Navigate to server directory:
```bash
cd server
```

2. Install Go dependencies:
```bash
go mod download
```

3. Install Python dependencies:
```bash
pip install tensorflow numpy nltk
```

4. Download NLTK data:
```bash
python -c "import nltk; nltk.download('stopwords'); nltk.download('wordnet')"
```

5. Run the server:
```bash
go run main.go
```

### Frontend Setup

The frontend is static HTML/CSS/JS and will be served by the Go backend automatically.

## 📡 API Endpoints

### `POST /api/predict`

Classify an email as spam or not spam.

**Request:**
```json
{
  "email_text": "Congratulations! You've won $1000000! Click here now!!!"
}
```

**Response:**
```json
{
  "is_spam": true,
  "confidence": 98.76,
  "label": "Spam"
}
```

### `GET /api/health`

Health check endpoint for monitoring.

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2025-12-04T10:30:00Z"
}
```

### Example Usage

```bash
# Using curl
curl -X POST http://localhost:8080/api/predict \
  -H "Content-Type: application/json" \
  -d '{"email_text": "Meeting scheduled for tomorrow at 10 AM"}'

# Using Python requests
import requests
response = requests.post('http://localhost:8080/api/predict',
    json={'email_text': 'Your email text here'})
print(response.json())
```

## Project Structure

```
project_n/
├── server/
│   ├── main.go              # Go backend server
│   ├── predict.py           # Python inference script
│   ├── EmailClassifier.keras # Trained model
│   ├── tokenizer_email.pickle # Tokenizer
│   ├── go.mod               # Go dependencies
│   └── Dockerfile           # Docker configuration
├── frontend/
│   ├── index.html           # Main HTML file
│   ├── styles.css           # Styling
│   └── script.js            # Frontend logic
├── model/
│   └── Spam-email.ipynb     # Model training notebook
├── docker-compose.yml       # Docker Compose configuration
└── README.md                # This file
```

## 🧠 Model Architecture

Our spam classifier uses a sophisticated **Bidirectional LSTM** neural network that reads email text in both forward and backward directions for better context understanding.

### Model Specifications

| Metric | Value |
|--------|-------|
| **Accuracy** | 97% on test set |
| **Architecture** | Bidirectional LSTM |
| **Framework** | TensorFlow 2.14 / Keras |
| **Dataset** | Kaggle Spam Email Dataset |
| **Vocabulary Size** | 5,000 words |
| **Max Sequence Length** | 1,000 tokens |

### Network Layers

```python
Embedding Layer (5000 → 1000)
    ↓
Bidirectional LSTM (128 units, dropout=0.5)
    ↓
Bidirectional LSTM (64 units, dropout=0.5)
    ↓
Dense Layer (2 units, softmax activation)
```

### Preprocessing Pipeline

1. **Text Cleaning** - Remove subject prefixes
2. **Stop Word Removal** - Filter common English words
3. **Lemmatization** - Reduce words to base forms
4. **Tokenization** - Convert to numerical sequences
5. **Padding** - Normalize sequence lengths

📊 **Want to train your own model?** Check out `model/Spam-email.ipynb` for the complete training pipeline!

## 💻 Development

### Testing the API

```bash
# Test spam email
curl -X POST http://localhost:8080/api/predict \
  -H "Content-Type: application/json" \
  -d '{"email_text": "WINNER! Click here to claim your FREE prize NOW!!!"}'

# Test legitimate email
curl -X POST http://localhost:8080/api/predict \
  -H "Content-Type: application/json" \
  -d '{"email_text": "Hi team, the project meeting is scheduled for 3 PM today."}'
```

### Building Docker Image Manually

```bash
docker build -t spam-classifier -f server/Dockerfile .
docker run -p 8080:8080 spam-classifier
```

### Running Without Docker

```bash
# Terminal 1: Start the backend
cd server
go run main.go

# Terminal 2: Access via browser or API
# http://localhost:8080
```

## 🐛 Troubleshooting

<details>
<summary>Docker build fails with "input/output error"</summary>

```bash
# Clean Docker system
docker system prune -a --volumes -f

# Rebuild
docker-compose up --build
```
</details>

<details>
<summary>Module import errors in Python</summary>

Make sure you're using Python 3.10 and have installed dependencies:
```bash
pip install -r model/requirements.txt
```
</details>

<details>
<summary>Port 8080 already in use</summary>

```bash
# Change port in docker-compose.yml
ports:
  - "3000:8080"  # Use port 3000 instead
```
</details>

## 📸 Screenshots

<div align="center">

### Web Interface
![Web Interface](https://via.placeholder.com/800x400?text=Spam+Classifier+Interface)

### API Response
![API Response](https://via.placeholder.com/800x300?text=API+JSON+Response)

</div>

## 🎯 Future Enhancements

- [ ] Add support for multiple languages
- [ ] Implement real-time email monitoring
- [ ] Add confidence threshold settings
- [ ] Create mobile app (React Native)
- [ ] Add batch processing for multiple emails
- [ ] Implement A/B testing for model improvements
- [ ] Add email attachment scanning

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/SahithiKokkula/Spam-Email-Classifier/issues).

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 👥 Authors

**Sahithi Kokkula** - [GitHub](https://github.com/SahithiKokkula)

## 🙏 Acknowledgments

- Dataset: [Kaggle Spam Email Dataset](https://www.kaggle.com/datasets/jackksoncsie/spam-email-dataset)
- Inspired by real-world email security challenges
- Built with ❤️ for the open-source community

## ⭐ Show Your Support

Give a ⭐️ if this project helped you!

---

<div align="center">

**[⬆ Back to Top](#-spam-email-classifier)**

Made with 💙 by Sahithi Kokkula

</div>

