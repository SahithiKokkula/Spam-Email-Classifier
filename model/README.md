# Spam Email Classification Model

This directory contains the training notebook and related files for the spam email classification model.

## Model Overview

The model is a deep learning LSTM (Long Short-Term Memory) neural network designed to classify emails as spam or not spam.

## Architecture

- **Type**: Bidirectional LSTM
- **Layers**:
  - Embedding layer (5000 vocabulary, 1000 dimensions)
  - Bidirectional LSTM (128 units, dropout 0.5)
  - Bidirectional LSTM (64 units, dropout 0.5)
  - Dense layer (2 units, softmax activation)

## Dataset

- **Source**: [Kaggle Spam Email Dataset](https://www.kaggle.com/datasets/jackksoncsie/spam-email-dataset)
- **Size**: 5,728 emails
- **Distribution**:
  - Not Spam (0): 4,360 emails
  - Spam (1): 1,368 emails

## Preprocessing

1. Remove "Subject:" prefix from emails
2. Remove stop words
3. Lemmatize words
4. Tokenize and pad sequences to length 1000
5. Limit vocabulary to top 5000 most frequent words

## Training

- **Framework**: TensorFlow/Keras
- **Optimizer**: Adam
- **Loss**: Categorical Crossentropy
- **Metrics**: Accuracy
- **Train/Test Split**: 70/30
- **Epochs**: 1 (model converged quickly)

## Performance

- **Accuracy**: 97%
- **Precision (Spam)**: 0.98
- **Recall (Spam)**: 0.92
- **F1-Score (Spam)**: 0.95

## Files

- `Spam-email.ipynb`: Jupyter notebook with complete training pipeline
- `emails.csv`: Dataset file
- `requirements.txt`: Python dependencies

## Saved Model Files

The trained model and tokenizer are saved in the `server/` directory:
- `EmailClassifier.keras`: Saved Keras model
- `tokenizer_email.pickle`: Saved tokenizer for text preprocessing

## Requirements

```
pandas>=2.1.0
numpy>=1.26.0
scikit-learn>=1.3.0
tensorflow>=2.14.0
keras>=2.14.0
nltk>=3.8.1
h5py>=3.10.0
```

## Usage

The model is used by the backend server for real-time predictions. See the main `README.md` for deployment instructions.

