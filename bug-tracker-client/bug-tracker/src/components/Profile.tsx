import styled from 'styled-components';
import { useState, useContext, useEffect } from 'react';
import { UserContext } from '../context/userContext';
import jwtDecode, { JwtPayload } from 'jwt-decode';
interface ProfileDetails {
  name?: string;
  email?: string;
  img?: string;
  bio?: string;
  contact?: string;
  admin?: boolean;
}

const profiles = [
  {
    name: 'Rhys Halpin',
    img: '/profileTest.png',
    description:
      'I am a full stack engineer, Current project lead of Bug-Tracker and currently learning Golang and React',
    contact: 'rhyshalpin@hotmail.co.uk',
    admin: true,
  },
];

const Profile: React.FC<ProfileDetails> = () => {
  const { state, actions } = useContext(UserContext);
  const [loading, setLoading] = useState(true);
  console.log(state.user.client_id);
  console.log(state);

  const [data, setData] = useState<ProfileDetails>();

  useEffect(() => {
    fetchProfile();
  }, []);

  const fetchProfile = async () => {
    try {
      const token: string = state.user;
      console.log(token);
      const decoded = jwtDecode<JwtPayload>(token);
      console.log(decoded.client_id);
      const res = await fetch('http://localhost:8000/apiv1/auth/profile', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          userObject: decoded.client_id,
        }),
      });
      const data = await res.json();
      setLoading(false);
      console.log(data);
      setData(data);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <Container>
      <h2>Profile</h2>
      {loading ? (
        'loading...'
      ) : (
        <Contents>
          <img src={profiles[0].img} alt="profileImg" />
          <p>{data?.name}</p>
          <p>{data?.bio}</p>
          <p>Role: {profiles[0].admin === true ? 'Admin' : 'Engineer'}</p>
          <p>{data?.email}</p>
          <button>Edit Profile</button>
        </Contents>
      )}
    </Container>
  );
};

export default Profile;

const Container = styled.div`
  border: 3px solid pink;
  width: 30%;
  word-wrap: break-word;
  height: 50%;

  img {
    width: 100px;
    height: 100px;
    border-radius: 50%;
    border: solid 3px;
  }
`;
const Contents = styled.div``;
