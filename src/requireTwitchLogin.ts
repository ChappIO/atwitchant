
export async function requireTwitchLogin() {
    const newToken = new URLSearchParams(window.location.hash.substring(1)).get('access_token');

    if(newToken) {
        await fetch('/auth/twitch', {
            method: 'POST',
            body: JSON.stringify({
                token: newToken
            }),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        window.location.hash = '';
    } else {
        const response = await fetch('/auth/twitch');

        switch (response.status) {
            case 404:
                const body = await response.json();
                // we don't have a status code just yet...
                window.location.assign(body.redirectUri);
                break;
            case 200:
                // the user has already logged in :D
                break;
        }
    }
}
