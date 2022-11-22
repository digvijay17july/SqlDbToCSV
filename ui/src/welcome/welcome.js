import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Container from 'react-bootstrap/Container';
import Alert from 'react-bootstrap/Alert';
import { useState } from 'react';
import { useNavigate } from "react-router-dom";

function Welcome(props) {
  const [host, setHost] = useState("localhost")
  const [port, setPort] = useState(5432)
  const [user, setUser] = useState("postgres")
  const [password, setPassword] = useState("admin")
  const [dbname, setDbName] = useState("dvdrental")
  const [rowsPerFile, setRowsPerFile] = useState(100)

  let navigate = useNavigate(); 
  let handleSubmit = async (e) => {
    e.preventDefault();
    try {
      let res = await fetch("http://localhost:4000/api/loadConfig", {
        method: "POST",
        body: JSON.stringify({
          "host": host,
          "port": port,
          "user": user,
          "password": password,
          "dbname": dbname,
          "rowsPerFile": rowsPerFile
        }),
      });
      let resJson = await res.json();
      if (res.status === 200) {

        console.log("Config created successfully");
        navigate('/migration');
      } else {
        console.log("Some error occured");
      }
    } catch (err) {
      console.log(err);
    }
  };
  return (
    <Alert variant="primary" className="d-flex align-items-center justify-content-center  min-vh-100">

      <Container className="position-absolute top-50 start-50 translate-middle shadow p-3 mb-5 bg-body rounded min-vh-100">
        <Alert.Heading>Settings for DB Connection!</Alert.Heading>
        <p>
          <Form onSubmit={handleSubmit}>
            <Form.Group className="mb-3" controlId="formBasicUsernameID">
              <Form.Label >Username</Form.Label>
              <Form.Control type="text" placeholder={user} />
              <Form.Text className="text-muted" value={user} onChange={(e) => setUser(e.target.value)}>
              Please Enter the DB Username
            </Form.Text>
          </Form.Group>

          <Form.Group className="mb-3" controlId="formBasicPassword">
            <Form.Label>Password</Form.Label>
            <Form.Control type="password" placeholder={password} />
              <Form.Text className="text-muted" value={password} onChange={(e) => setPassword(e.target.value)} >
              Please Enter the DB Password
            </Form.Text>
          </Form.Group>
          <Form.Group className="mb-3" controlId="formBasicDbIP">
            <Form.Label>DB IP ADDRESS(HOST)</Form.Label>
            <Form.Control type="text" placeholder={host} />
              <Form.Text className="text-muted" value={host} onChange={(e) => setHost(e.target.value)}>
              Please Enter the DB IP ADDRESS
            </Form.Text>
          </Form.Group>
          <Form.Group className="mb-3" controlId="formBasicDbPort">
            <Form.Label>PORT</Form.Label>
            <Form.Control type="text" placeholder={port} />
              <Form.Text className="text-muted" value={port} onChange={(e) => setPort(e.target.value)}>
              Please Enter the Port
            </Form.Text>
          </Form.Group>
          <Form.Group className="mb-3" controlId="formBasicDbName">
            <Form.Label>DB Name</Form.Label>
            <Form.Control type="text" placeholder={dbname} />
              <Form.Text className="text-muted" value={dbname} onChange={(e) => setDbName(e.target.value)}>
              Please Enter the DB Name
            </Form.Text>
          </Form.Group>
          <Form.Group className="mb-3" controlId="formBasicRowsPerFile">
            <Form.Label>Rows Per File</Form.Label>
            <Form.Control type="text" placeholder={rowsPerFile} />
              <Form.Text className="text-muted" value={rowsPerFile} onChange={(e) => setRowsPerFile(e.target.value)}>
              Please Enter the Rows Per File
            </Form.Text>
          </Form.Group>
          <Button variant="primary" type="submit">
            NEXT
          </Button>
        </Form>
      </p>
    </Container>
    </Alert >
  )
}

export default Welcome