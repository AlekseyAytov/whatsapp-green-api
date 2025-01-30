async function getSettings() {
    const idInstance = document.getElementById('idInstance').value;
    const apiTokenInstance = document.getElementById('apiTokenInstance').value;
    const responseOutput = document.getElementById('responseOutput');

    try {
        responseOutput.textContent = ""
        // const response = await fetch('http://178.208.80.96:8080/settings', {
        const response = await fetch('http://127.0.0.1:8080/settings', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                idInstance: idInstance,
                apiTokenInstance: apiTokenInstance,
            }),
        });

        // Проверяем, удачный ли ответ
        if (response.ok) {
            const data = await response.json();
            var str = JSON.stringify(data, null, 1);
            responseOutput.textContent = str; // Выводим ответ
        } else {
            responseOutput.textContent = 'Ошибка при получении настроек';
        }
    } catch (error) {
        responseOutput.textContent = 'Ошибка сети: ' + error.message;
    }
}

async function getStateInstance() {
    const idInstance = document.getElementById('idInstance').value;
    const apiTokenInstance = document.getElementById('apiTokenInstance').value;
    const responseOutput = document.getElementById('responseOutput');

    try {
        responseOutput.textContent = ""
        // const response = await fetch('http://178.208.80.96:8080/state', {
        const response = await fetch('http://127.0.0.1:8080/state', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                idInstance: idInstance,
                apiTokenInstance: apiTokenInstance,
            }),
        });

        // Проверяем, удачный ли ответ
        if (response.ok) {
            const data = await response.json();
            var str = JSON.stringify(data, null, 1);
            responseOutput.textContent = str; // Выводим ответ
        } else {
            responseOutput.textContent = 'Ошибка при получении настроек';
        }
    } catch (error) {
        responseOutput.textContent = 'Ошибка сети: ' + error.message;
    }
}

async function sendMessage() {
    const idInstance = document.getElementById('idInstance').value;
    const apiTokenInstance = document.getElementById('apiTokenInstance').value;
    const numberForSend = document.getElementById('messageInput').value;
    const textForSend = document.getElementById('messageText').value;
    const responseOutput = document.getElementById('responseOutput');

    try {
        responseOutput.textContent = ""
        // const response = await fetch('http://178.208.80.96:8080/send', {
        const response = await fetch('http://127.0.0.1:8080/send', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                idInstance: idInstance,
                apiTokenInstance: apiTokenInstance,
                number: numberForSend,
                text: textForSend,
            }),
        });

        // Проверяем, удачный ли ответ
        if (response.ok) {
            const data = await response.json();
            var str = JSON.stringify(data, null, 1);
            responseOutput.textContent = str; // Выводим ответ
        } else {
            responseOutput.textContent = 'Ошибка при получении настроек';
        }
    } catch (error) {
        responseOutput.textContent = 'Ошибка сети: ' + error.message;
    }
}

async function sendFileByUrl() {
    const idInstance = document.getElementById('idInstance').value;
    const apiTokenInstance = document.getElementById('apiTokenInstance').value;
    const numberForSend = document.getElementById('numberFile').value;
    const fileForSend = document.getElementById('fileInput').value;
    const responseOutput = document.getElementById('responseOutput');

    try {
        responseOutput.textContent = ""
        // const response = await fetch('http://178.208.80.96:8080/file', {
        const response = await fetch('http://127.0.0.1:8080/file', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                idInstance: idInstance,
                apiTokenInstance: apiTokenInstance,
                number: numberForSend,
                file: fileForSend,
            }),
        });

        // Проверяем, удачный ли ответ
        if (response.ok) {
            const data = await response.json();
            var str = JSON.stringify(data, null, 1);
            responseOutput.textContent = str; // Выводим ответ
        } else {
            responseOutput.textContent = 'Ошибка при получении настроек';
        }
    } catch (error) {
        responseOutput.textContent = 'Ошибка сети: ' + error.message;
    }
}