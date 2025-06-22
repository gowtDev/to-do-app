
<button onClick={handleLogout} className="mb-4 bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700">
  Logout
</button>
const handleLogout = () => {
  localStorage.removeItem('token');
  window.location.href = '/login';
};

<form onSubmit={handleCreateTodo} className="mb-6 bg-white p-4 shadow rounded">
  <h3 className="text-xl mb-2">Add Todo</h3>
  <input name="title" placeholder="Title" className="w-full mb-2 p-2 border rounded" required />
  <input name="description" placeholder="Description" className="w-full mb-2 p-2 border rounded" required />
  <button className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">Add</button>
</form>
const handleCreateTodo = async (e) => {
  e.preventDefault();
  const token = localStorage.getItem('token');
  const title = e.target.title.value;
  const description = e.target.description.value;

  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/todos`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token,
    },
    body: JSON.stringify({ title, description }),
  });

  if (res.ok) {
    e.target.reset();
    fetchTodos(); // refresh todo list
  }
};


function isTokenExpired(token) {
  try {
    const payload = JSON.parse(atob(token.split('.')[1]));
    const expiry = payload.exp;
    return Date.now() >= expiry * 1000;
  } catch (e) {
    return true; // Invalid token
  }
}


useEffect(() => {
  const token = localStorage.getItem('token');
  if (!token || isTokenExpired(token)) {
    localStorage.removeItem('token');
    alert('Session expired. Please login again.');
    window.location.href = '/login';
  } else {
    fetchTodos();
  }
}, []);
