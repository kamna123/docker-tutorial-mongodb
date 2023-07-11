document.getElementById('addUserForm').addEventListener('submit', addUser);

function addUser(e) {
  e.preventDefault();
  const username = document.getElementById('username').value;
  const email = document.getElementById('email').value;
  const user = { username, email };

  fetch('/users', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(user)
  })
    .then(response => response.json())
    .then(data => {
      console.log('User added:', data);
      document.getElementById('addUserForm').reset();
    fetchUserDetails();
    })
    .catch(error => {
      console.log('Error:', error);
    });
}

function fetchUserDetails() {
    fetch('/users')
      .then(response => {
        if (response.ok) {
          return response.json();
        } else {
          throw new Error('Failed to fetch user details');
        }
      })
      .then(users => {
        const userDetails = document.getElementById('userDetails');
        userDetails.innerHTML = '';
        users.forEach(user => {
          userDetails.innerHTML += `
            <div>
              <p><strong>ID:</strong> ${user.id}</p>
              <p><strong>Username:</strong> ${user.username}</p>
              <p><strong>Email:</strong> ${user.email}</p>
            </div>
            <hr>
          `;
        });
      })
      .catch(error => {
        console.log('Error:', error);
      });
  }
  
