import styled from 'styled-components';

interface ProfileDetails {
  name?: string;
  img?: string;
  description?: string;
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
  return (
    <Container>
      <h2>Profile</h2>

      <img src={profiles[0].img} alt="profileImg" />
      <p>{profiles[0].name}</p>
      <p>{profiles[0].description}</p>
      <p>Role: {profiles[0].admin === true ? 'Admin' : 'Engineer'}</p>
      <p>{profiles[0].contact}</p>
      <button>Edit Profile</button>
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
