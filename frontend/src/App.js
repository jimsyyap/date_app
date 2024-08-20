import React, { useState, useEffect } from 'react';

const App = () => {
    const [data, setData] = useState([]);

    useEffect(() => {
        fetch('/api/data')
            .then(response => response.json())
            .then(data => setData(data))
            .catch(error => console.error('Error fetching data:', error));
    }, []);

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4">
            <h1 className="text-4xl font-bold text-blue-600 mb-8">Hello, World!</h1>
            <div className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
                <h2 className="text-2xl font-semibold mb-4">Data from PostgreSQL:</h2>
                <ul className="list-disc pl-5">
                    {data.map(item => (
                        <li key={item.id} className="mb-2">
                            {item.name}
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    );
};

export default App;
