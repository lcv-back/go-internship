import React from 'react';

const Signup = () => {
  return (
    <div className='wrapper'>
      <form action="">
        <div>
        <input type="text" name="username" placeholder='Username' required />
          </div>
        <div>
          <input type="password" name="password" placeholder='Password' required />
        </div>
        <div>
          <input type="email" name="email" placeholder='Email' required />
        </div>
        <button type="submit">Sign up</button>
    </form>
    </div>
  );
}

export default Signup; 
