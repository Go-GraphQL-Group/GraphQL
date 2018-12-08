# Schema Types

<details>
  <summary><strong>Table of Contents</strong></summary>

  * [Query](#query)
  * [Objects](#objects)
    * [BrowseQuery](#browsequery)
    * [Film](#film)
    * [FilmConnection](#filmconnection)
    * [FilmEdge](#filmedge)
    * [LookupQuery](#lookupquery)
    * [PageInfo](#pageinfo)
    * [People](#people)
    * [PeopleConnection](#peopleconnection)
    * [PeopleEdge](#peopleedge)
    * [Planet](#planet)
    * [PlanetConnection](#planetconnection)
    * [PlanetEdge](#planetedge)
    * [SearchQuery](#searchquery)
    * [Specie](#specie)
    * [SpecieConnection](#specieconnection)
    * [SpecieEdge](#specieedge)
    * [Starship](#starship)
    * [StarshipConnection](#starshipconnection)
    * [StarshipEdge](#starshipedge)
    * [Vehicle](#vehicle)
    * [VehicleConnection](#vehicleconnection)
    * [VehicleEdge](#vehicleedge)
  * [Enums](#enums)
    * [Gender](#gender)
    * [VehicleClass](#vehicleclass)
  * [Scalars](#scalars)
    * [Boolean](#boolean)
    * [ID](#id)
    * [Int](#int)
    * [String](#string)

</details>

## Query
The query root, from which multiple types of requests can be made.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>lookup</strong></td>
<td valign="top"><a href="#lookupquery">LookupQuery</a></td>
<td>

Perform a lookup of an entity by its ID.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>browse</strong></td>
<td valign="top"><a href="#browsequery">BrowseQuery</a></td>
<td>

Browse all Star Wars entities.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>search</strong></td>
<td valign="top"><a href="#searchquery">SearchQuery</a></td>
<td>

Search for entities using Lucene query syntax.

</td>
</tr>
</tbody>
</table>

## Objects

### BrowseQuery

A query for all Star Wars entities.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>peoples</strong></td>
<td valign="top"><a href="#peopleconnection">PeopleConnection</a>!</td>
<td>

Browse people entities.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>films</strong></td>
<td valign="top"><a href="#filmconnection">FilmConnection</a>!</td>
<td>

Browse film entities.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>starships</strong></td>
<td valign="top"><a href="#starshipconnection">StarshipConnection</a>!</td>
<td>

Browse starship entities.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>vehicles</strong></td>
<td valign="top"><a href="#vehicleconnection">VehicleConnection</a>!</td>
<td>

Browse vehicle entities.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>species</strong></td>
<td valign="top"><a href="#specieconnection">SpecieConnection</a>!</td>
<td>

Browse specie entities.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>planets</strong></td>
<td valign="top"><a href="#planetconnection">PlanetConnection</a>!</td>
<td>

Browse planet entities.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
</tbody>
</table>

### Film

A Film resource is a single film.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The id of this film.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>title</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The title of this film.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>episode_id</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The episode number of this film.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>opening_crawl</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The opening paragraphs at the beginning of this film.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>director</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The name of the director of this film.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>producer</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The name(s) of the producer(s) of this film. Comma separated.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>release_date</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The ISO 8601 date format of film release at original creator country.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>species</strong></td>
<td valign="top">[<a href="#specie">Specie</a>]</td>
<td>

An array of species resources that are in this film.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>starships</strong></td>
<td valign="top">[<a href="#starship">Starship</a>]</td>
<td>

An array of starship resources that are in this film.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>vehicles</strong></td>
<td valign="top">[<a href="#vehicle">Vehicle</a>]</td>
<td>

An array of vehicle resources that are in this film.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>characters</strong></td>
<td valign="top">[<a href="#vehicle">Vehicle</a>]</td>
<td>

An array of Vehicle resources that are in this film.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>planets</strong></td>
<td valign="top">[<a href="#planet">Planet</a>]</td>
<td>

An array of planet resources that are in this film.

</td>
</tr>
</tbody>
</table>

### FilmConnection

A connection to a list of items.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>pageInfo</strong></td>
<td valign="top"><a href="#pageinfo">PageInfo</a>!</td>
<td>

Information to aid in pagination.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>edges</strong></td>
<td valign="top">[<a href="#filmedge">FilmEdge</a>!]</td>
<td>

A list of edges.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>nodes</strong></td>
<td valign="top">[<a href="#film">Film</a>!]</td>
<td>

A list of nodes in the connection (without going through the `edges` field).

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>totalCount</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

A count of the total number of items in this connection, ignoring pagination.

</td>
</tr>
</tbody>
</table>

### FilmEdge

An edge in a connection.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>node</strong></td>
<td valign="top"><a href="#film">Film</a></td>
<td>

The item at the end of the edge.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cursor</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

A cursor for use in pagination.

</td>
</tr>
</tbody>
</table>

### LookupQuery

A lookup of an individual Star Wars entity by its ID.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>people</strong></td>
<td valign="top"><a href="#people">People</a></td>
<td>

Look up a specific people by its ID.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The ID of the entity.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>film</strong></td>
<td valign="top"><a href="#film">Film</a></td>
<td>

Look up a specific film by its ID.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The ID of the entity.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>starship</strong></td>
<td valign="top"><a href="#starship">Starship</a></td>
<td>

Look up a specific starship by its ID.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The ID of the entity.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>vehicle</strong></td>
<td valign="top"><a href="#vehicle">Vehicle</a></td>
<td>

Look up a specific vehicle by its ID.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The ID of the entity.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>specie</strong></td>
<td valign="top"><a href="#specie">Specie</a></td>
<td>

Look up a specific specie by its ID.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The ID of the entity.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>planet</strong></td>
<td valign="top"><a href="#planet">Planet</a></td>
<td>

Look up a specific planet by its ID.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The ID of the entity.

</td>
</tr>
</tbody>
</table>

### PageInfo

Information about pagination in a connection.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>hasNextPage</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td>

When paginating forwards, are there more items?

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>hasPreviousPage</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td>

When paginating backwards, are there more items?

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>startCursor</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

When paginating backwards, the cursor to continue.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>endCursor</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

When paginating forwards, the cursor to continue.

</td>
</tr>
</tbody>
</table>

### People

A People resource is an individual person or character within the Star Wars universe.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The id of this person.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The name of this person.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>birth_year</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The birth year of the person, using the in-universe standard of BBY or ABY - Before the Battle of Yavin or After the Battle of Yavin. The Battle of Yavin is a battle that occurs at the end of Star Wars episode IV: A New Hope.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>eye_color</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The eye color of this person. Will be "unknown" if not known or "n/a" if the person does not have an eye.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>gender</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The gender of this person. Either "Male", "Female" or "unknown", "n/a" if the person does not have a gender.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>hair_color</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The hair color of this person. Will be "unknown" if not known or "n/a" if the person does not have hair.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>height</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The height of the person in centimeters.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>mass</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The mass of the person in kilograms.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>skin_color</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The skin color of this person.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>homeworld</strong></td>
<td valign="top"><a href="#planet">Planet</a></td>
<td>

A planet resource, a planet that this person was born on or inhabits.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>films</strong></td>
<td valign="top">[<a href="#film">Film</a>]</td>
<td>

An array of film resources that this person has been in.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>species</strong></td>
<td valign="top">[<a href="#specie">Specie</a>]</td>
<td>

An array of specie resources that this person belongs to.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>starships</strong></td>
<td valign="top">[<a href="#starship">Starship</a>]</td>
<td>

An array of starship resources that this person has piloted.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>vehicles</strong></td>
<td valign="top">[<a href="#vehicle">Vehicle</a>]</td>
<td>

An array of vehicle resources that this person has piloted.

</td>
</tr>
</tbody>
</table>

### PeopleConnection

A connection to a list of items.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>pageInfo</strong></td>
<td valign="top"><a href="#pageinfo">PageInfo</a>!</td>
<td>

Information to aid in pagination.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>edges</strong></td>
<td valign="top">[<a href="#peopleedge">PeopleEdge</a>!]</td>
<td>

A list of edges.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>nodes</strong></td>
<td valign="top">[<a href="#people">People</a>!]</td>
<td>

A list of nodes in the connection (without going through the `edges` field).

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>totalCount</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

A count of the total number of items in this connection, ignoring pagination.

</td>
</tr>
</tbody>
</table>

### PeopleEdge

An edge in a connection.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>node</strong></td>
<td valign="top"><a href="#people">People</a></td>
<td>

The item at the end of the edge.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cursor</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

A cursor for use in pagination.

</td>
</tr>
</tbody>
</table>

### Planet

A Planet resource is a large mass, planet or planetoid in the Star Wars Universe, at the time of 0 ABY.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The id of this planet.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The name of this planet.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>diameter</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The diameter of this planet in kilometers.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>rotation_period</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The number of standard hours it takes for this planet to complete a single rotation on its axis.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>orbital_period</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The number of standard days it takes for this planet to complete a single orbit of its local star.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>gravity</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

A number denoting the gravity of this planet, where "1" is normal or 1 standard G. "2" is twice or 2 standard Gs. "0.5" is half or 0.5 standard Gs.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>population</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The average population of sentient beings inhabiting this planet.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>climate</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The climate of this planet. Comma separated if diverse.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>terrain</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The terrain of this planet. Comma separated if diverse.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>surface_water</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The percentage of the planet surface that is naturally occurring water or bodies of water.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>residents</strong></td>
<td valign="top">[<a href="#vehicle">Vehicle</a>]</td>
<td>

An array of Vehicle resources that live on this planet.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>films</strong></td>
<td valign="top">[<a href="#film">Film</a>]</td>
<td>

An array of film resources that this planet has appeared in.

</td>
</tr>
</tbody>
</table>

### PlanetConnection

A connection to a list of items.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>pageInfo</strong></td>
<td valign="top"><a href="#pageinfo">PageInfo</a>!</td>
<td>

Information to aid in pagination.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>edges</strong></td>
<td valign="top">[<a href="#planetedge">PlanetEdge</a>!]</td>
<td>

A list of edges.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>nodes</strong></td>
<td valign="top">[<a href="#planet">Planet</a>!]</td>
<td>

A list of nodes in the connection (without going through the `edges` field).

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>totalCount</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

A count of the total number of items in this connection, ignoring pagination.

</td>
</tr>
</tbody>
</table>

### PlanetEdge

An edge in a connection.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>node</strong></td>
<td valign="top"><a href="#planet">Planet</a></td>
<td>

The item at the end of the edge.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cursor</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

A cursor for use in pagination.

</td>
</tr>
</tbody>
</table>

### SearchQuery

A search for Star Wars entities using Lucene query syntax.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>peoples</strong></td>
<td valign="top"><a href="#peopleconnection">PeopleConnection</a></td>
<td>

Search for people entities matching the given query.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">search</td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The search field for name, in Lucene search syntax.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>films</strong></td>
<td valign="top"><a href="#filmconnection">FilmConnection</a></td>
<td>

Search for film entities matching the given query.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">search</td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The search field for title, in Lucene search syntax.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>starships</strong></td>
<td valign="top"><a href="#starshipconnection">StarshipConnection</a></td>
<td>

Search for starship entities matching the given query.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">search</td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The search field for name or model, in Lucene search syntax.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>vehicles</strong></td>
<td valign="top"><a href="#vehicleconnection">VehicleConnection</a></td>
<td>

Search for vehicle entities matching the given query.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">search</td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The search field for name or model, in Lucene search syntax.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>species</strong></td>
<td valign="top"><a href="#specieconnection">SpecieConnection</a></td>
<td>

Search for specie entities matching the given query.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">search</td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The search field for name, in Lucene search syntax.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>planets</strong></td>
<td valign="top"><a href="#planetconnection">PlanetConnection</a></td>
<td>

Search for planet entities matching the given query.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">search</td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The search field for name, in Lucene search syntax.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">first</td>
<td valign="top"><a href="#int">Int</a></td>
<td>

The number of entities in the connection.

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">after</td>
<td valign="top"><a href="#id">ID</a></td>
<td>

The connection follows by.

</td>
</tr>
</tbody>
</table>

### Specie

A Species resource is a type of person or character within the Star Wars Universe.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The id of this specie.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The name of this species.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>classification</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The classification of this species, such as "mammal" or "reptile".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>designation</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The designation of this species, such as "sentient".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>average_height</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The average height of this species in centimeters.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>average_lifespan</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The average lifespan of this species in years.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>eye_colors</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

A comma-separated string of common eye colors for this species, "none" if this species does not typically have eyes.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>hair_colors</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

A comma-separated string of common hair colors for this species, "none" if this species does not typically have hair.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>skin_colors</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

A comma-separated string of common skin colors for this species, "none" if this species does not typically have skin.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>language</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The language commonly spoken by this species.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>homeworld</strong></td>
<td valign="top"><a href="#planet">Planet</a></td>
<td>

A planet resource, a planet that this species originates from.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>Vehicle</strong></td>
<td valign="top">[<a href="#vehicle">Vehicle</a>]</td>
<td>

An array of Vehicle resources that are a part of this species.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>films</strong></td>
<td valign="top">[<a href="#film">Film</a>]</td>
<td>

An array of film resources that this species has appeared in.

</td>
</tr>
</tbody>
</table>

### SpecieConnection

A connection to a list of items.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>pageInfo</strong></td>
<td valign="top"><a href="#pageinfo">PageInfo</a>!</td>
<td>

Information to aid in pagination.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>edges</strong></td>
<td valign="top">[<a href="#specieedge">SpecieEdge</a>!]</td>
<td>

A list of edges.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>nodes</strong></td>
<td valign="top">[<a href="#specie">Specie</a>!]</td>
<td>

A list of nodes in the connection (without going through the `edges` field).

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>totalCount</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

A count of the total number of items in this connection, ignoring pagination.

</td>
</tr>
</tbody>
</table>

### SpecieEdge

An edge in a connection.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>node</strong></td>
<td valign="top"><a href="#specie">Specie</a></td>
<td>

The item at the end of the edge.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cursor</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

A cursor for use in pagination.

</td>
</tr>
</tbody>
</table>

### Starship

A Starship resource is a single transport craft that has hyperdrive capability.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The id of this starship.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The name of this starship. The common name, such as "Death Star".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>model</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The model or official name of this starship. Such as "T-65 X-wing" or "DS-1 Orbital Battle Station".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>starship_class</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The class of this starship, such as "Starfighter" or "Deep Space Mobile Battlestation"

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>manufacturer</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The manufacturer of this starship. Comma separated if more than one.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cost_in_credits</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The cost of this starship new, in galactic credits.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>length</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The length of this starship in meters.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>crew</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The number of personnel needed to run or pilot this starship.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>passengers</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The number of non-essential Vehicle this starship can transport.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>max_atmosphering_speed</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The maximum speed of this starship in the atmosphere. "N/A" if this starship is incapable of atmospheric flight.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>hyperdrive_rating</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The class of this starships hyperdrive.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>MGLT</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The Maximum number of Megalights this starship can travel in a standard hour. A "Megalight" is a standard unit of distance and has never been defined before within the Star Wars universe. This figure is only really useful for measuring the difference in speed of starships. We can assume it is similar to AU, the distance between our Sun (Sol) and Earth.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cargo_capacity</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The maximum number of kilograms that this starship can transport.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>consumables</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The maximum length of time that this starship can provide consumables for its entire crew without having to resupply.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>films</strong></td>
<td valign="top">[<a href="#film">Film</a>]</td>
<td>

An array of film resources that this starship has appeared in.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>pilots</strong></td>
<td valign="top">[<a href="#vehicle">Vehicle</a>]</td>
<td>

An array of Vehicle resources that this starship has been piloted by.

</td>
</tr>
</tbody>
</table>

### StarshipConnection

A connection to a list of items.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>pageInfo</strong></td>
<td valign="top"><a href="#pageinfo">PageInfo</a>!</td>
<td>

Information to aid in pagination.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>edges</strong></td>
<td valign="top">[<a href="#starshipedge">StarshipEdge</a>!]</td>
<td>

A list of edges.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>nodes</strong></td>
<td valign="top">[<a href="#starship">Starship</a>!]</td>
<td>

A list of nodes in the connection (without going through the `edges` field).

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>totalCount</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

A count of the total number of items in this connection, ignoring pagination.

</td>
</tr>
</tbody>
</table>

### StarshipEdge

An edge in a connection.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>node</strong></td>
<td valign="top"><a href="#starship">Starship</a></td>
<td>

The item at the end of the edge.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cursor</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

A cursor for use in pagination.

</td>
</tr>
</tbody>
</table>

### Vehicle

A Vehicle resource is a single transport craft that does not have hyperdrive capability.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

The id of this vehicle.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The name of this vehicle. The common name, such as "Sand Crawler" or "Speeder bike".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>model</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The model or official name of this vehicle. Such as "All-Terrain Attack Transport".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>vehicle_class</strong></td>
<td valign="top"><a href="#vehicleclass">VehicleClass</a></td>
<td>

The class of this vehicle, such as "Wheeled" or "Repulsorcraft".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>manufacturer</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The manufacturer of this vehicle. Comma separated if more than one.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>length</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The length of this vehicle in meters.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cost_in_credits</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The cost of this vehicle new, in Galactic Credits.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>crew</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The number of personnel needed to run or pilot this vehicle.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>passengers</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The number of non-essential Vehicle this vehicle can transport.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>max_atmosphering_speed</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The maximum speed of this vehicle in the atmosphere.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cargo_capacity</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The maximum number of kilograms that this vehicle can transport.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>consumables</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The maximum length of time that this vehicle can provide consumables for its entire crew without having to resupply.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>films</strong></td>
<td valign="top">[<a href="#film">Film</a>]</td>
<td>

An array of film resources that this vehicle has appeared in.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>pilots</strong></td>
<td valign="top">[<a href="#vehicle">Vehicle</a>]</td>
<td>

An array of Vehicle resources that this vehicle has been piloted by.

</td>
</tr>
</tbody>
</table>

### VehicleConnection

A connection to a list of items.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>pageInfo</strong></td>
<td valign="top"><a href="#pageinfo">PageInfo</a>!</td>
<td>

Information to aid in pagination.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>edges</strong></td>
<td valign="top">[<a href="#vehicleedge">VehicleEdge</a>!]</td>
<td>

A list of edges.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>nodes</strong></td>
<td valign="top">[<a href="#vehicle">Vehicle</a>!]</td>
<td>

A list of nodes in the connection (without going through the `edges` field).

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>totalCount</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

A count of the total number of items in this connection, ignoring pagination.

</td>
</tr>
</tbody>
</table>

### VehicleEdge

An edge in a connection.

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>node</strong></td>
<td valign="top"><a href="#vehicle">Vehicle</a></td>
<td>

The item at the end of the edge.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cursor</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

A cursor for use in pagination.

</td>
</tr>
</tbody>
</table>

## Enums

### Gender

The genders of people in the Star Wars universe.

<table>
<thead>
<th align="left">Value</th>
<th align="left">Description</th>
</thead>
<tbody>
<tr>
<td valign="top"><strong>NA</strong></td>
<td>

Does not have a gender.

</td>
</tr>
<tr>
<td valign="top"><strong>MALE</strong></td>
<td>

Male.

</td>
</tr>
<tr>
<td valign="top"><strong>FEMALE</strong></td>
<td>

Female.

</td>
</tr>
</tbody>
</table>

### VehicleClass

The class of vehicle in the Star Wars universe.

<table>
<thead>
<th align="left">Value</th>
<th align="left">Description</th>
</thead>
<tbody>
<tr>
<td valign="top"><strong>WHEELED</strong></td>
<td>

Wheeled.

</td>
</tr>
<tr>
<td valign="top"><strong>REPULSOCRAFT</strong></td>
<td>

Repulsorcraft.

</td>
</tr>
</tbody>
</table>

## Scalars

### Boolean

The `Boolean` scalar type represents `true` or `false`.

### ID

The `ID` scalar type represents a unique identifier, often used to refetch an object or as key for a cache. The ID type appears in a JSON response as a String; however, it is not intended to be human-readable. When expected as an input type, any string (such as `"4"`) or integer (such as `4`) input value will be accepted as an ID.

### Int

The `Int` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1. 

### String

The `String` scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.

