import styled from 'styled-components';
import Feed from './Feed';
import Profile from './Profile';
import SideBar from './SideBar';

const Dashboard: React.FC = () => {
  return (
    <Container>
      <div>WELCOME TO THE DASHBOARD</div>
      <DashContainer>
        <SideBar />
        <Main>
          <Feed />
          <Profile />
          Dashboard main
        </Main>
      </DashContainer>
    </Container>
  );
};

export default Dashboard;

const Container = styled.div`
  height: 100%;
  border: solid 3px black;
  align-items: stretch;
`;

const Main = styled.main`
  border: solid 3px blue;
  width: 90%;
  display: flex;
  flex-direction: row;
  align-items: stretch;
`;

const DashContainer = styled.div`
  height: 85vh;
  border: solid 3px black;
  display: flex;
  flex-direction: row;
  align-items: stretch;
`;
