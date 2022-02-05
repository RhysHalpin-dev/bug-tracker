import styled from 'styled-components';

const Header: React.FC = () => {
  return (
    <Nav>
      <NavMenu>
        <span>BUG TRACKER</span>
      </NavMenu>
    </Nav>
  );
};

export default Header;

const Nav = styled.nav`
  height: 70px;
  border: solid 3px black;
  display: flex;
  align-items: center;
`;

const NavMenu = styled.div`
  display: flex;
  flex: 1;
  margin-left: 25px;
  align-items: center;
  max-width: 100%;
  span {
    font-size: 13px;
    font-weight: 650;
    letter-spacing: 1.42px;
    position: relative;

    margin-left: 5px;
  }
`;
