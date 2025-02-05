import { useEffect, useState } from "react";

function App() {
    const [message, setMessage] = useState("");

    useEffect(() => {
        fetch("http://localhost:8080/")
            .then((res) => res.text())
            .then((data) => setMessage(data))
            .catch((err) => console.error("Error fetching data:", err));
    }, []);

    return (
        <div>
            <h1>Collaborative Project Management Platform</h1>
            <p>{message}</p>
        </div>
    );
}

export default App;
