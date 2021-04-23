function sendRequest() {
    (async () => {
        const httpMethod = Array.from(document.getElementsByName('typeSelector')).filter(element => element.checked)[0].value
        let url = '/organizations'
        let id = document.getElementById('org-id').value
        let name = document.getElementById('org-name').value
        let trial = document.getElementById('org-trial').checked
        let requestBody = {
            name: name,
            trial: trial
        }
        if (httpMethod === 'GET') {
            url += `?id=${id}`
            requestBody = undefined
        }
        if (httpMethod === 'DELETE') {
            url += `/${id}`
        }
        const response = await fetch(url, {
            headers: { "Content-Type": "application/json" },
            method: httpMethod,
            body: JSON.stringify(requestBody)
        })
        let outputLocation = document.getElementById('result')
        let responseString = ""
        if (response.ok) {
            try {
                const data = await response.json()
                for (const [property, value] of Object.entries(data)) {
                    responseString += `${property}: ${value}\n`
                }
            } catch (e) {
                responseString = "Success. This command does not display any output."
            }
        } else {
            responseString += "Error! "
            switch (response.status) {
                case 404:
                    responseString += "The organization you entered does not exist."
                    break
                case 400:
                    responseString += "Check that the formatting of your form entries are correct and try again."
                    break
                default:
                    responseString += "An unknown error occurred."
            }
        }
        outputLocation.innerText = responseString
    })()
}

let idOnlyBoxes = [document.getElementById('get'), document.getElementById('delete')]

idOnlyBoxes.forEach(box => box.addEventListener('change', function () {
    if (this.checked) {
        document.getElementById('field-id').style.display = 'block'
        document.getElementById('field-name').style.display = 'none'
        document.getElementById('field-trial').style.display = 'none'
    }
}))

document.getElementById('post').addEventListener('change', function () {
    if (this.checked) {
        document.getElementById('field-id').style.display = 'none'
        document.getElementById('field-name').style.display = 'block'
        document.getElementById('field-trial').style.display = 'block'
    }
})

document.getElementById('patch').addEventListener('change', function () {
    if (this.checked) {
        document.getElementById('field-id').style.display = 'block'
        document.getElementById('field-name').style.display = 'block'
        document.getElementById('field-trial').style.display = 'block'
    }
})

document.getElementById('field-id').style.display = 'none' // the first instruction selected is 'create' so ID should be none