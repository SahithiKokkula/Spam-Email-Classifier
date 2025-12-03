const API_URL = '/api/predict';

function showLoading() {
    document.getElementById('loadingSection').classList.remove('hidden');
    document.getElementById('resultSection').classList.add('hidden');
    document.getElementById('errorSection').classList.add('hidden');
    document.getElementById('predictBtn').disabled = true;
}

function hideLoading() {
    document.getElementById('loadingSection').classList.add('hidden');
    document.getElementById('predictBtn').disabled = false;
}

function showResult(result) {
    const resultSection = document.getElementById('resultSection');
    const resultCard = document.getElementById('resultCard');
    const resultLabel = document.getElementById('resultLabel');
    const resultConfidence = document.getElementById('resultConfidence');

    resultCard.className = 'result-card ' + (result.is_spam ? 'spam' : 'not-spam');
    resultLabel.textContent = result.label;
    resultConfidence.textContent = `Confidence: ${result.confidence}%`;

    resultSection.classList.remove('hidden');
    hideLoading();
}

function showError(message) {
    const errorSection = document.getElementById('errorSection');
    const errorMessage = document.getElementById('errorMessage');
    
    errorMessage.textContent = message;
    errorSection.classList.remove('hidden');
    hideLoading();
}

async function predictEmail() {
    const emailText = document.getElementById('emailInput').value.trim();

    if (!emailText) {
        showError('Please enter email text');
        return;
    }

    showLoading();

    try {
        const response = await fetch(API_URL, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email_text: emailText }),
        });

        const data = await response.json();

        if (data.error) {
            showError(data.error);
        } else {
            showResult(data);
        }
    } catch (error) {
        showError('Failed to connect to server. Please make sure the server is running.');
        console.error('Error:', error);
    }
}

document.getElementById('emailInput').addEventListener('keydown', function(e) {
    if (e.ctrlKey && e.key === 'Enter') {
        predictEmail();
    }
});

