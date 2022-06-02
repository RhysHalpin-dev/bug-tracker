import { Link } from 'react-router-dom';
import styled from 'styled-components';

const SideBar: React.FC = () => {
  return (
    <Container>
      sideBar
      <h1>WELCOME USER!</h1>
      <ul>
        <li>
          <Link to="dashIndex">Dashboard</Link>
        </li>
        <li>
          <Link to="dashTickets">My Tickets</Link>
        </li>
        <li>
          <Link to="dashProjects">My Projects</Link>
        </li>
        <li>
          <Link to="history">History</Link>
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
