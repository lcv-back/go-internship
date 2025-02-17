import React from 'react';

const Login = () => {
  return (
    <div className='wrapper'>
      <form action="">
        <div>
        <input type="text" name="username" placeholder='Username' required />
          </div>
        <div>
          <input type="password" name="password" placeholder='Password' required />
        </div>
        <button type="submit">Login</button>
    </form>
    </div>
  );
}

export default Login; 
