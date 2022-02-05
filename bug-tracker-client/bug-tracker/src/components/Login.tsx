import styled from 'styled-components';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login: React.FC = () => {
  const [user, setUser] = useState('');
  const [pass, setPass] = useState('');
  let navigate = useNavigate();

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    console.log(user + ' ' + pass);
    navigate('./dashboard', { replace: true });
    return (
      <div>
        ${user}${pass}
      </div>
    );
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
