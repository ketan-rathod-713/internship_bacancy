import React, { useState } from 'react';
import './global.css';

const LoginForm = () => {
  const [formData, setFormData] = useState({
    userId: '',
    password: ''
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // You can add form submission logic here
    console.log(formData);

    // create post request to server
     fetch({url: "http://localhost:8080/api/user", method: "POST", headers: {"Content-Type": "application/json"}, body: JSON.stringify(formData)}).then(response => response.json()).then(data => console.log(data))
  };

  return (
    <div className='signuppage'>
      <form className="signup-form" onSubmit={handleSubmit}>
      <div>
        <label htmlFor="userId">User Id</label>
        <input type="text" id="userId" name="userId" value={formData.userId} onChange={handleChange} required />
      </div>
      <div>
        <label htmlFor="password">Password</label>
        <input type="password" id="password" name="password" value={formData.password} onChange={handleChange} required />
      </div>
      <div className='signnupbtn'>
        <button>Submit</button>
      </div>
    </form>
    </div>
  );
};

export default LoginForm;
