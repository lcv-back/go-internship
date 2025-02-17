import React, { useState } from 'react';
import Login from './pages/Login';
import Signup from './pages/Signup';
import './App.css';

function App() {
  const [isLogin, setIsLogin] = useState(true); // Chúng ta sử dụng trạng thái này để xác định đang hiển thị form nào

  return (
    <div className="App">
      <div className="form-container">
        <div className="form-box">
          <h1>{isLogin ? 'Login' : 'Sign Up'}</h1>
          {isLogin ? (
            <>
              <Login />
              <div className="signup-link">
                <p>Don't have an account? <a href="#" onClick={() => setIsLogin(false)}>Sign Up</a></p>
              </div>
            </>
          ) : (
            <>
              <Signup />
              <div className="login-link">
                <p>Already have an account? <a href="#" onClick={() => setIsLogin(true)}>Login</a></p>
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
