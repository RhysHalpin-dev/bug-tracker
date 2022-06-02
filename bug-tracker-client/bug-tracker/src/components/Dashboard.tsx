import styled from 'styled-components';
import Feed from './Feed';
import Profile from './Profile';
import SideBar from './SideBar';
import { Routes, Route, Link, Outlet } from 'react-router-dom';

const Dashboard: React.FC = () => {
  return (
    <Container>
      <div>WELCOME TO THE DASHBOARD</div>
      <DashContainer>
        <SideBar />
        <Outlet />
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

const DashContainer = styled.div`
  height: 85vh;
  border: solid 3px black;
  display: flex;
  flex-direction: row;
  align-items: stretch;
`;
