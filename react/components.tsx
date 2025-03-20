import React from 'react';
import { useState, useEffect } from "react";


function Counter() {
    const [count, setCount] = useState(0);

    useEffect(() => {
        const dataContainer = document.getElementById("data");
        if (!dataContainer) {
            return;
        }

        const data = dataContainer.getAttribute("x-data");
        if (data) {
            const storedCount = data.toString();
            setCount(parseInt(storedCount));
        }
    }, []);

    return (
        <div className="flex flex-col items-center p-4 border rounded-lg shadow-md">
            <h1 className="text-2xl font-bold">Counter: {count}</h1>
            <button
                onClick={async () => {
                    const result = await fetch("/count", {
                        method: "GET",
                        headers: {
                            "Content-Type": "application/json",
                        },
                    });

                    const data = await result.json();
                    setCount(data.count);
                }}
                className="mt-2 px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
            >
                Increment
            </button>
        </div>
    );
}

export const Body = () => (
    <div>
        <Counter />
    </div>
);
