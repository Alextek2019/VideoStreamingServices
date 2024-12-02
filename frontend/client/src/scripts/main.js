// Redirect to register page
function redirectToRegister() {
  window.location.href = "register.html";
}

// Redirect to login page
function redirectToLogin() {
  window.location.href = "login.html";
}

// Handle login
async function login() {
  const login = document.getElementById('login-username').value;
  const password = document.getElementById('login-password').value;

  if (!login || !password) {
    alert('Please fill in all fields.');
    return;
  }

  try {
    const response = await fetch('http://localhost:8000/api/v1/user', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ username: login, password })
    });

    if (response.ok) {
      const { token } = await response.json();
      localStorage.setItem('token', token);
      alert('Login successful!');
      window.location.href = "home.html"; // Redirect to home page
    } else {
      alert('Login failed. Please check your credentials.');
    }
  } catch (error) {
    alert('An error occurred during login.');
  }
}

// Handle registration
async function register() {
  const login = document.getElementById('register-username').value;
  const password = document.getElementById('register-password').value;

  if (!login || !password) {
    alert('Please fill in all fields.');
    return;
  }

  try {
    const response = await fetch('http://localhost:8000/api/v1/user', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({login: login, password: password})
    });

    if (response.ok) {
      alert('Registration successful! Please log in.');
      window.location.href = "login.html";
    } else {
      alert('Registration failed.');
    }
  } catch (error) {
    alert('An error occurred during registration.');
  }
}

// Logout
function logout() {
  localStorage.removeItem('token');
  window.location.href = "login.html";
}
