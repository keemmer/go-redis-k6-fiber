import http from 'k6/http';

export let options = {
    vus: 5,
    duration: '5s'
}

export default function () {
    // http.get('http://localhost:8000/hello')
    http.get('http://host.docker.internal:8000/hello')
}

// docker-compose run --rm k6 run /scripts/test.js