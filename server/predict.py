import sys
import json
import pickle
import numpy as np
import tensorflow as tf
from tensorflow import keras
from tensorflow.keras.preprocessing.sequence import pad_sequences
import nltk
from nltk.corpus import stopwords
from nltk.tokenize import WhitespaceTokenizer
from nltk.stem import WordNetLemmatizer

import os
nltk_data_path = '/usr/local/nltk_data'
if os.path.exists(nltk_data_path):
    nltk.data.path.append(nltk_data_path)

try:
    stop_words = set(stopwords.words('english'))
except LookupError:
    nltk.download('stopwords')
    stop_words = set(stopwords.words('english'))

try:
    lemmatizer = WordNetLemmatizer()
except LookupError:
    nltk.download('wordnet')
    lemmatizer = WordNetLemmatizer()

w_tokenizer = WhitespaceTokenizer()

def rem_subject(text):
    if text.startswith('Subject:'):
        return text[9:].strip()
    return text.strip()

def remove_stop_words(sentence):
    if isinstance(sentence, str):
        words = sentence.split()
        filtered_words = [word for word in words if word.lower() not in stop_words]
        return ' '.join(filtered_words)
    return ""

def lemmatize_text(text):
    if isinstance(text, str):
        st = ""
        for w in w_tokenizer.tokenize(text):
            st = st + lemmatizer.lemmatize(w) + " "
        return st.strip()
    return ""

def preprocess_text(text):
    text = rem_subject(text)
    text = remove_stop_words(text)
    text = lemmatize_text(text)
    return text

def predict(email_text):
    max_sequence_length = 1000
    max_word = 5000
    
    with open('tokenizer_email.pickle', 'rb') as handle:
        tokenizer = pickle.load(handle)
    
    model = keras.models.load_model('EmailClassifier.keras')
    
    processed_text = preprocess_text(email_text)
    seq = tokenizer.texts_to_sequences([processed_text])
    padded_seq = pad_sequences(seq, maxlen=max_sequence_length)
    
    prediction = model.predict(padded_seq, verbose=0)
    predicted_class = int(np.around(prediction[0], decimals=0).argmax())
    confidence = float(prediction[0][predicted_class])
    
    result = {
        "is_spam": bool(predicted_class),
        "confidence": round(confidence * 100, 2),
        "label": "Spam" if predicted_class else "Not Spam"
    }
    
    return result

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print(json.dumps({"error": "No email text provided"}))
        sys.exit(1)
    
    email_text = sys.argv[1]
    result = predict(email_text)
    print(json.dumps(result))

