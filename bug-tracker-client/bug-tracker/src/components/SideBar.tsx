import styled from 'styled-components';

const SideBar: React.FC = () => {
  return (
    <Container>
      sideBar
      <h1>WELCOME USER!</h1>
      <ul>
        <li>
          <a href="http://localhost:3000/dashboard">Dashboard</a>
        </li>
        <li>
          <a href="#">My Tickets</a>
        </li>
        <li>
          <a href="#">My Projects</a>
        </li>
        <li>
          <a href="#">History</a>
        </li>
      </ul>
    </Container>
  );
};

export default SideBar;

const Container = styled.nav`
  width: 15rem;
  border: solid 3px red;
  text-align: center;

  ul {
    margin: 0;

    border: solid black 3px;
    list-style: none;
    li {
      margin: 10px;
      margin-right: 40%;
      a {
        text-decoration: none;
      }
    }
  }
`;
