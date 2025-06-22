import { useState } from 'react';

export default function RegisterPage() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  const handleRegister = async (e) => {
    e.preventDefault();

    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/iam/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, email, password }),
    });

    const data = await res.json();

    if (res.ok) {
      localStorage.setItem('token', data.token);
      setMessage('✅ Registration successful!');
    } else {
      setMessage('❌ Registration failed');
    }
  };

  return (
    <form onSubmit={handleRegister} className="bg-white p-6 rounded shadow-md w-96 mx-auto mt-20">
      <h2 className="text-2xl mb-4">Register</h2>
      <input type="text" placeholder="Name" className="w-full mb-2 p-2 border rounded"
        value={name} onChange={(e) => setName(e.target.value)} required />
      <input type="email" placeholder="Email" className="w-full mb-2 p-2 border rounded"
        value={email} onChange={(e) => setEmail(e.target.value)} required />
      <input type="password" placeholder="Password" className="w-full mb-4 p-2 border rounded"
        value={password} onChange={(e) => setPassword(e.target.value)} required />
      <button type="submit" className="w-full bg-green-600 text-white py-2 rounded hover:bg-green-700">
        Register
      </button>
      {message && <p className="mt-4 text-center text-red-500">{message}</p>}
    </form>
  );
}
