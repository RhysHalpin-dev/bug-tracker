import styled from 'styled-components';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login: React.FC = () => {
  const [user, setUser] = useState('');
  const [pass, setPass] = useState('');
  const [loginError, setLogginError] = useState('');
  let navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log(user + ' ' + pass);
    // POST login credentials to API login endpoint
    try {
      const res = await fetch('http://localhost:8000/apiv1/auth/Login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: user,
          password: pass,
        }),
      });
      const data = await res.json();

      console.log(data);

      if (res.status === 200) {
        navigate('./dashboard', { replace: true });
      } else {
        setLogginError('Incorrect login credentials, please try again');
      }
    } catch (err) {
      console.log(err);
    }
  };
  return (
    <LoginContainer>
      <LoginForm onSubmit={handleSubmit}>
        <h1>Bug Tracker Login</h1>
        <div>
          <label>Email:</label>
          <input
            type="email"
            name="userName"
            placeholder="Username"
            required
            onChange={(e) => setUser(e.target.value)}
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            name="password"
            placeholder="Password"
            required
            onChange={(e) => setPass(e.target.value)}
          />
        </div>
        <p id="errorMessage">{loginError}</p>
        <input type="submit" value="Login" />
        <p>
          Forgot your <a href="">password?</a>
        </p>
        <p>
          Create an account?<a href="">Sign Up?</a>
        </p>
        <p>
          sign in as <a href="">Demo User?</a>
        </p>
      </LoginForm>
    </LoginContainer>
  );
};

export default Login;

const LoginContainer = styled.main`
  display: flex;
  align-items: center; // vertical
  justify-content: center; // horizontal
  border: 3px solid black;
  height: 90vh;
`;

const LoginForm = styled.form`
  display: flex;
  border-radius: 1rem;
  justify-content: center; // horizontal
  flex-direction: column;
  border: 3px solid black;
  padding: 0.5em;
  font-weight: 500;
  h1 {
    font-weight: 800;
    letter-spacing: 1.42px;
    position: relative;

    margin-left: 5px;
  }
  p {
    margin: 0px;
  }
  #errorMessage {
    color: red;
    padding: 5px;
    border: solid 1px blue;
    border-radius: 15px;
    background-color: grey;
    opacity: 1;
    transition: all 5s;
  }
  div {
    display: flex;
    flex-direction: row;
    margin-bottom: 10px;
  }
  label {
    flex-basis: 130px;
  }
  input[type='email'] {
  }
  input[type='password'] {
  }
  input[type='submit'] {
    padding: 10px;
    font-size: 1em;
    background-color: blueviolet;
    color: white;
  }
`;
