<!DOCTYPE html>
<html>
<head>
	<title>Create an article</title>
</head>
<body>
		<form action="/articles" method="post" >
			<label for="id">ID:</label><br>
  			<input type="text" id="id" name="id" ><br>
  			<label for="title">Title:</label><br>
  			<input type="text" id="title" name="title" ><br>
  			<label for="stitle">SubTitle:</label><br>
  			<input type="text" id="stitle" name="stitle"><br>
  			<label for="content">Content:</label><br>
  			<input type="text" id="content" name="content"><br>
  			<label for="timestamp">TimeStamp:</label><br>
  			<input type="time" name="timestamp"><br><br>

 			<input type="submit" value="Submit">
		</form>
</body>
</html>