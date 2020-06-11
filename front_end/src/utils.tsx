export const request = (method: string, url: string, callback: (data: any) => void) => {
    const req = new XMLHttpRequest();
    req.open(method, url, true);

    if (!req) throw 'xhr creation failed';
    req.onerror = () => { throw 'xhr call failed'; }

    req.onload = () => {
        const data = JSON.parse(req.responseText);
        callback(data);
    }

    req.send();
}