


# Ressources API
Ce package centralise de multiples modules de ressource d'API.
Toutes les ressources impémentent entièrement ou partiellement l'interface "Ressource" du framework Gapi.
Le but de ce package est des fournir des ressources d'API typiques déja documentées et testées.
Toutes ces ressources peuvent être utilisées dans le framework Gapi par un simple import.

Framework Gapi: https://github.com/BoyerDamien/gapi
  

## Informations

### Version

1.0.0

### License

[MIT](http://opensource.org/licenses/MIT)

### Contact

 damienboyer45@gmail.com 

## Content negotiation

### URI Schemes
  * http
  * https

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  media

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/media | [create media](#create-media) |  |
| DELETE | /api/v1/media/{id} | [delete media](#delete-media) |  |
| DELETE | /api/v1/medias | [delete media list](#delete-media-list) |  |
| GET | /api/v1/medias | [media list](#media-list) |  |
| GET | /api/v1/media/{id} | [retrieve media](#retrieve-media) |  |
| PUT | /api/v1/media | [update media](#update-media) |  |
  


###  tag

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/tag | [create tag](#create-tag) |  |
| DELETE | /api/v1/tag/{id} | [delete tag](#delete-tag) |  |
| DELETE | /api/v1/tags | [delete tag list](#delete-tag-list) |  |
| GET | /api/v1/tag/{id} | [retrieve tag](#retrieve-tag) |  |
| GET | /api/v1/tags | [tag list](#tag-list) |  |
  


###  user

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/user | [create user](#create-user) |  |
| DELETE | /api/v1/user/{id} | [delete user](#delete-user) |  |
| DELETE | /api/v1/users | [delete user list](#delete-user-list) |  |
| GET | /api/v1/user/{id} | [retrieve user](#retrieve-user) |  |
| PUT | /api/v1/user | [update](#update) |  |
| GET | /api/v1/users | [user list](#user-list) |  |
  


## Paths

### <span id="create-media"></span> create media (*CreateMedia*)

```
POST /api/v1/media
```

Créé un nouveau média

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| media | `formData` | file | `io.ReadCloser` |  |  |  | Contenu du média |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-media-200) | OK | Retourne le média créé |  | [schema](#create-media-200-schema) |
| [400](#create-media-400) | Bad Request | StatusBadRequest |  | [schema](#create-media-400-schema) |
| [500](#create-media-500) | Internal Server Error | StatusInternalServerError |  | [schema](#create-media-500-schema) |
| [default](#create-media-default) | | Erreur |  | [schema](#create-media-default-schema) |

#### Responses


##### <span id="create-media-200"></span> 200 - Retourne le média créé
Status: OK

###### <span id="create-media-200-schema"></span> Schema
   
  

[Media](#media)

##### <span id="create-media-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="create-media-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="create-media-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="create-media-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="create-media-default"></span> Default Response
Erreur

###### <span id="create-media-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="create-tag"></span> create tag (*CreateTag*)

```
POST /api/v1/tag
```

Créé un nouveau tag

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tag | `body` | [Tag](#tag) | `models.Tag` | |  | | Données du tag |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-tag-200) | OK | Retourne le tag |  | [schema](#create-tag-200-schema) |
| [400](#create-tag-400) | Bad Request | StatusBadRequest |  | [schema](#create-tag-400-schema) |
| [500](#create-tag-500) | Internal Server Error | StatusInternalServerError |  | [schema](#create-tag-500-schema) |
| [default](#create-tag-default) | | Erreur |  | [schema](#create-tag-default-schema) |

#### Responses


##### <span id="create-tag-200"></span> 200 - Retourne le tag
Status: OK

###### <span id="create-tag-200-schema"></span> Schema
   
  

[Tag](#tag)

##### <span id="create-tag-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="create-tag-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="create-tag-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="create-tag-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="create-tag-default"></span> Default Response
Erreur

###### <span id="create-tag-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="create-user"></span> create user (*CreateUser*)

```
POST /api/v1/user
```

Créé un nouvel utilisateur

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user | `body` | [User](#user) | `models.User` | |  | | Données de l'utilisateur |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-user-200) | OK | Retourne l'utilisateur créé |  | [schema](#create-user-200-schema) |
| [400](#create-user-400) | Bad Request | StatusBadRequest |  | [schema](#create-user-400-schema) |
| [500](#create-user-500) | Internal Server Error | StatusInternalServerError |  | [schema](#create-user-500-schema) |
| [default](#create-user-default) | | Erreur |  | [schema](#create-user-default-schema) |

#### Responses


##### <span id="create-user-200"></span> 200 - Retourne l'utilisateur créé
Status: OK

###### <span id="create-user-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="create-user-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="create-user-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="create-user-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="create-user-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="create-user-default"></span> Default Response
Erreur

###### <span id="create-user-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="delete-media"></span> delete media (*DeleteMedia*)

```
DELETE /api/v1/media/{id}
```

Supprime un média existant

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | nom du média |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-media-200) | OK | Valide la suppression |  | [schema](#delete-media-200-schema) |
| [202](#delete-media-202) | Accepted | StatusAccepted |  | [schema](#delete-media-202-schema) |
| [500](#delete-media-500) | Internal Server Error | StatusInternalServerError |  | [schema](#delete-media-500-schema) |
| [default](#delete-media-default) | | Erreur |  | [schema](#delete-media-default-schema) |

#### Responses


##### <span id="delete-media-200"></span> 200 - Valide la suppression
Status: OK

###### <span id="delete-media-200-schema"></span> Schema

##### <span id="delete-media-202"></span> 202 - StatusAccepted
Status: Accepted

###### <span id="delete-media-202-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-media-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="delete-media-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-media-default"></span> Default Response
Erreur

###### <span id="delete-media-default-schema"></span> Schema
empty schema

### <span id="delete-media-list"></span> delete media list (*DeleteMediaList*)

```
DELETE /api/v1/medias
```

Supprime une liste de médias

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| names | `query` | []string | `[]string` |  | ✓ |  | Liste de noms |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-media-list-200) | OK | Valide la suppression |  | [schema](#delete-media-list-200-schema) |
| [202](#delete-media-list-202) | Accepted | StatusAccepted |  | [schema](#delete-media-list-202-schema) |
| [400](#delete-media-list-400) | Bad Request | StatusBadRequest |  | [schema](#delete-media-list-400-schema) |
| [500](#delete-media-list-500) | Internal Server Error | StatusInternalServerError |  | [schema](#delete-media-list-500-schema) |
| [default](#delete-media-list-default) | | Erreur |  | [schema](#delete-media-list-default-schema) |

#### Responses


##### <span id="delete-media-list-200"></span> 200 - Valide la suppression
Status: OK

###### <span id="delete-media-list-200-schema"></span> Schema

##### <span id="delete-media-list-202"></span> 202 - StatusAccepted
Status: Accepted

###### <span id="delete-media-list-202-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-media-list-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="delete-media-list-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-media-list-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="delete-media-list-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-media-list-default"></span> Default Response
Erreur

###### <span id="delete-media-list-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="delete-tag"></span> delete tag (*DeleteTag*)

```
DELETE /api/v1/tag/{id}
```

Supprime un tag existant

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | nom tu tag |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-tag-200) | OK | Valide la suppression |  | [schema](#delete-tag-200-schema) |
| [202](#delete-tag-202) | Accepted | StatusAccepted |  | [schema](#delete-tag-202-schema) |
| [500](#delete-tag-500) | Internal Server Error | StatusInternalServerError |  | [schema](#delete-tag-500-schema) |
| [default](#delete-tag-default) | | Erreur |  | [schema](#delete-tag-default-schema) |

#### Responses


##### <span id="delete-tag-200"></span> 200 - Valide la suppression
Status: OK

###### <span id="delete-tag-200-schema"></span> Schema

##### <span id="delete-tag-202"></span> 202 - StatusAccepted
Status: Accepted

###### <span id="delete-tag-202-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-tag-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="delete-tag-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-tag-default"></span> Default Response
Erreur

###### <span id="delete-tag-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="delete-tag-list"></span> delete tag list (*DeleteTagList*)

```
DELETE /api/v1/tags
```

Supprime une liste de tags

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| names | `query` | []string | `[]string` |  | ✓ |  | Liste de noms |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-tag-list-200) | OK | Valide la suppression |  | [schema](#delete-tag-list-200-schema) |
| [202](#delete-tag-list-202) | Accepted | StatusAccepted |  | [schema](#delete-tag-list-202-schema) |
| [400](#delete-tag-list-400) | Bad Request | StatusBadRequest |  | [schema](#delete-tag-list-400-schema) |
| [500](#delete-tag-list-500) | Internal Server Error | StatusInternalServerError |  | [schema](#delete-tag-list-500-schema) |
| [default](#delete-tag-list-default) | | Erreur |  | [schema](#delete-tag-list-default-schema) |

#### Responses


##### <span id="delete-tag-list-200"></span> 200 - Valide la suppression
Status: OK

###### <span id="delete-tag-list-200-schema"></span> Schema

##### <span id="delete-tag-list-202"></span> 202 - StatusAccepted
Status: Accepted

###### <span id="delete-tag-list-202-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-tag-list-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="delete-tag-list-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-tag-list-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="delete-tag-list-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-tag-list-default"></span> Default Response
Erreur

###### <span id="delete-tag-list-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="delete-user"></span> delete user (*DeleteUser*)

```
DELETE /api/v1/user/{id}
```

Supprime un utilisateur existant

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | email de l'utilisateur |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-user-200) | OK | Valide la suppression |  | [schema](#delete-user-200-schema) |
| [202](#delete-user-202) | Accepted | StatusAccepted |  | [schema](#delete-user-202-schema) |
| [500](#delete-user-500) | Internal Server Error | StatusInternalServerError |  | [schema](#delete-user-500-schema) |
| [default](#delete-user-default) | | Erreur |  | [schema](#delete-user-default-schema) |

#### Responses


##### <span id="delete-user-200"></span> 200 - Valide la suppression
Status: OK

###### <span id="delete-user-200-schema"></span> Schema

##### <span id="delete-user-202"></span> 202 - StatusAccepted
Status: Accepted

###### <span id="delete-user-202-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-user-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="delete-user-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-user-default"></span> Default Response
Erreur

###### <span id="delete-user-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="delete-user-list"></span> delete user list (*DeleteUserList*)

```
DELETE /api/v1/users
```

Supprime une liste d'utilisateurs

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| emails | `query` | []string | `[]string` |  | ✓ |  | Liste de mails |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-user-list-200) | OK | Valide la suppression |  | [schema](#delete-user-list-200-schema) |
| [202](#delete-user-list-202) | Accepted | StatusAccepted |  | [schema](#delete-user-list-202-schema) |
| [400](#delete-user-list-400) | Bad Request | StatusBadRequest |  | [schema](#delete-user-list-400-schema) |
| [500](#delete-user-list-500) | Internal Server Error | StatusInternalServerError |  | [schema](#delete-user-list-500-schema) |
| [default](#delete-user-list-default) | | Erreur |  | [schema](#delete-user-list-default-schema) |

#### Responses


##### <span id="delete-user-list-200"></span> 200 - Valide la suppression
Status: OK

###### <span id="delete-user-list-200-schema"></span> Schema

##### <span id="delete-user-list-202"></span> 202 - StatusAccepted
Status: Accepted

###### <span id="delete-user-list-202-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-user-list-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="delete-user-list-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-user-list-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="delete-user-list-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="delete-user-list-default"></span> Default Response
Erreur

###### <span id="delete-user-list-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="media-list"></span> media list (*MediaList*)

```
GET /api/v1/medias
```

Retourne des informations détaillées sur une liste de médias

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| limit | `query` | number | `float64` |  |  |  | Limite le nombre de résultats au nombre passé en paramètre |
| offset | `query` | number | `float64` |  |  |  | Filtre les résultats a partir de l'index passé en paramètre |
| orderBy | `query` | string | `string` |  |  |  | Permet de trier les résultats par champs |
| status | `query` | string | `string` |  |  |  | Permet le filtre par status |
| type | `query` | string | `string` |  |  |  | Permet le filtre par mime type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#media-list-200) | OK | Retourne une liste de médias |  | [schema](#media-list-200-schema) |
| [400](#media-list-400) | Bad Request | StatusBadRequest |  | [schema](#media-list-400-schema) |
| [404](#media-list-404) | Not Found | StatusNotFound |  | [schema](#media-list-404-schema) |
| [500](#media-list-500) | Internal Server Error | StatusInternalServerError |  | [schema](#media-list-500-schema) |
| [default](#media-list-default) | | Erreur |  | [schema](#media-list-default-schema) |

#### Responses


##### <span id="media-list-200"></span> 200 - Retourne une liste de médias
Status: OK

###### <span id="media-list-200-schema"></span> Schema
   
  

[][User](#user)

##### <span id="media-list-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="media-list-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="media-list-404"></span> 404 - StatusNotFound
Status: Not Found

###### <span id="media-list-404-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="media-list-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="media-list-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="media-list-default"></span> Default Response
Erreur

###### <span id="media-list-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="retrieve-media"></span> retrieve media (*RetrieveMedia*)

```
GET /api/v1/media/{id}
```

Retourne des informations détaillées sur un média

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | le nom du média |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#retrieve-media-200) | OK | Retourne un média |  | [schema](#retrieve-media-200-schema) |
| [400](#retrieve-media-400) | Bad Request | StatusBadRequest |  | [schema](#retrieve-media-400-schema) |
| [404](#retrieve-media-404) | Not Found | StatusNotFound |  | [schema](#retrieve-media-404-schema) |
| [500](#retrieve-media-500) | Internal Server Error | StatusInternalServerError |  | [schema](#retrieve-media-500-schema) |
| [default](#retrieve-media-default) | | Erreur |  | [schema](#retrieve-media-default-schema) |

#### Responses


##### <span id="retrieve-media-200"></span> 200 - Retourne un média
Status: OK

###### <span id="retrieve-media-200-schema"></span> Schema
   
  

[Media](#media)

##### <span id="retrieve-media-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="retrieve-media-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="retrieve-media-404"></span> 404 - StatusNotFound
Status: Not Found

###### <span id="retrieve-media-404-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="retrieve-media-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="retrieve-media-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="retrieve-media-default"></span> Default Response
Erreur

###### <span id="retrieve-media-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="retrieve-tag"></span> retrieve tag (*RetrieveTag*)

```
GET /api/v1/tag/{id}
```

Retourne des informations détaillées sur un tag

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | nom du tag |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#retrieve-tag-200) | OK | Retourne un tag |  | [schema](#retrieve-tag-200-schema) |
| [400](#retrieve-tag-400) | Bad Request | StatusBadRequest |  | [schema](#retrieve-tag-400-schema) |
| [404](#retrieve-tag-404) | Not Found | StatusNotFound |  | [schema](#retrieve-tag-404-schema) |
| [500](#retrieve-tag-500) | Internal Server Error | StatusInternalServerError |  | [schema](#retrieve-tag-500-schema) |
| [default](#retrieve-tag-default) | | Erreur |  | [schema](#retrieve-tag-default-schema) |

#### Responses


##### <span id="retrieve-tag-200"></span> 200 - Retourne un tag
Status: OK

###### <span id="retrieve-tag-200-schema"></span> Schema
   
  

[Tag](#tag)

##### <span id="retrieve-tag-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="retrieve-tag-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="retrieve-tag-404"></span> 404 - StatusNotFound
Status: Not Found

###### <span id="retrieve-tag-404-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="retrieve-tag-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="retrieve-tag-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="retrieve-tag-default"></span> Default Response
Erreur

###### <span id="retrieve-tag-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="retrieve-user"></span> retrieve user (*RetrieveUser*)

```
GET /api/v1/user/{id}
```

Retourne des informations détaillées sur un utilisateur

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | l'email de l'utilisateur |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#retrieve-user-200) | OK | Retourne un utilisateur |  | [schema](#retrieve-user-200-schema) |
| [400](#retrieve-user-400) | Bad Request | StatusBadRequest |  | [schema](#retrieve-user-400-schema) |
| [404](#retrieve-user-404) | Not Found | StatusNotFound |  | [schema](#retrieve-user-404-schema) |
| [500](#retrieve-user-500) | Internal Server Error | StatusInternalServerError |  | [schema](#retrieve-user-500-schema) |
| [default](#retrieve-user-default) | | Erreur |  | [schema](#retrieve-user-default-schema) |

#### Responses


##### <span id="retrieve-user-200"></span> 200 - Retourne un utilisateur
Status: OK

###### <span id="retrieve-user-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="retrieve-user-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="retrieve-user-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="retrieve-user-404"></span> 404 - StatusNotFound
Status: Not Found

###### <span id="retrieve-user-404-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="retrieve-user-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="retrieve-user-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="retrieve-user-default"></span> Default Response
Erreur

###### <span id="retrieve-user-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="tag-list"></span> tag list (*TagList*)

```
GET /api/v1/tags
```

Retourne des informations détaillées sur une liste de tags

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| limit | `query` | number | `float64` |  |  |  | Limite le nombre de résultats au nombre passé en paramètre |
| offset | `query` | number | `float64` |  |  |  | Filtre les résultats a partir de l'index passé en paramètre |
| orderBy | `query` | string | `string` |  |  |  | Permet de trier les résultats par champs |
| toFind | `query` | string | `string` |  |  |  | Permet le filtre par nom. Retourne les tags pour lesquels le nom contient la chaîne de caractères à rechercher |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#tag-list-200) | OK | Retourne une liste de tags |  | [schema](#tag-list-200-schema) |
| [400](#tag-list-400) | Bad Request | StatusBadRequest |  | [schema](#tag-list-400-schema) |
| [404](#tag-list-404) | Not Found | StatusNotFound |  | [schema](#tag-list-404-schema) |
| [500](#tag-list-500) | Internal Server Error | StatusInternalServerError |  | [schema](#tag-list-500-schema) |
| [default](#tag-list-default) | | Erreur |  | [schema](#tag-list-default-schema) |

#### Responses


##### <span id="tag-list-200"></span> 200 - Retourne une liste de tags
Status: OK

###### <span id="tag-list-200-schema"></span> Schema
   
  

[][Tag](#tag)

##### <span id="tag-list-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="tag-list-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="tag-list-404"></span> 404 - StatusNotFound
Status: Not Found

###### <span id="tag-list-404-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="tag-list-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="tag-list-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="tag-list-default"></span> Default Response
Erreur

###### <span id="tag-list-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="update"></span> update (*Update*)

```
PUT /api/v1/user
```

Modifie un utilisateur existant

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user | `body` | [User](#user) | `models.User` | |  | | Données de l'utilisateur |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-200) | OK | Retourne l'utilisateur modifié |  | [schema](#update-200-schema) |
| [400](#update-400) | Bad Request | StatusBadRequest |  | [schema](#update-400-schema) |
| [500](#update-500) | Internal Server Error | StatusInternalServerError |  | [schema](#update-500-schema) |
| [default](#update-default) | | Erreur |  | [schema](#update-default-schema) |

#### Responses


##### <span id="update-200"></span> 200 - Retourne l'utilisateur modifié
Status: OK

###### <span id="update-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="update-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="update-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="update-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="update-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="update-default"></span> Default Response
Erreur

###### <span id="update-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="update-media"></span> update media (*UpdateMedia*)

```
PUT /api/v1/media
```

Modifie un média existant

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| media | `body` | [Media](#media) | `models.Media` | |  | | Données du média |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-media-200) | OK | Retourne le média modifié |  | [schema](#update-media-200-schema) |
| [400](#update-media-400) | Bad Request | StatusBadRequest |  | [schema](#update-media-400-schema) |
| [500](#update-media-500) | Internal Server Error | StatusInternalServerError |  | [schema](#update-media-500-schema) |
| [default](#update-media-default) | | Erreur |  | [schema](#update-media-default-schema) |

#### Responses


##### <span id="update-media-200"></span> 200 - Retourne le média modifié
Status: OK

###### <span id="update-media-200-schema"></span> Schema
   
  

[Media](#media)

##### <span id="update-media-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="update-media-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="update-media-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="update-media-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="update-media-default"></span> Default Response
Erreur

###### <span id="update-media-default-schema"></span> Schema

  

[ErrResponse](#err-response)

### <span id="user-list"></span> user list (*UserList*)

```
GET /api/v1/users
```

Retourne des informations détaillées sur une liste d'utilisateurs

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| limit | `query` | number | `float64` |  |  |  | Limite le nombre de résultats au nombre passé en paramètre |
| offset | `query` | number | `float64` |  |  |  | Filtre les résultats a partir de l'index passé en paramètre |
| orderBy | `query` | string | `string` |  |  |  | Permet de trier les résultats par champs |
| role | `query` | string | `string` |  |  |  | Permet le filtre par role |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#user-list-200) | OK | Retourne une liste d'utilisateurs |  | [schema](#user-list-200-schema) |
| [400](#user-list-400) | Bad Request | StatusBadRequest |  | [schema](#user-list-400-schema) |
| [404](#user-list-404) | Not Found | StatusNotFound |  | [schema](#user-list-404-schema) |
| [500](#user-list-500) | Internal Server Error | StatusInternalServerError |  | [schema](#user-list-500-schema) |
| [default](#user-list-default) | | Erreur |  | [schema](#user-list-default-schema) |

#### Responses


##### <span id="user-list-200"></span> 200 - Retourne une liste d'utilisateurs
Status: OK

###### <span id="user-list-200-schema"></span> Schema
   
  

[][User](#user)

##### <span id="user-list-400"></span> 400 - StatusBadRequest
Status: Bad Request

###### <span id="user-list-400-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="user-list-404"></span> 404 - StatusNotFound
Status: Not Found

###### <span id="user-list-404-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="user-list-500"></span> 500 - StatusInternalServerError
Status: Internal Server Error

###### <span id="user-list-500-schema"></span> Schema
   
  

[ErrResponse](#err-response)

##### <span id="user-list-default"></span> Default Response
Erreur

###### <span id="user-list-default-schema"></span> Schema

  

[ErrResponse](#err-response)

## Models

### <span id="err-response"></span> ErrResponse


> ErrResponse
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` | ✓ | | Message d'erreur | `error` |



### <span id="media"></span> Media


> Media
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` | ✓ | | Nom du média | `image.png` |
| Size | uint64 (formatted integer)| `uint64` | ✓ | | Taille du fichier en ko | `140` |
| Status | string| `string` | ✓ | | Status du l'image | `open` |
| Type | string| `string` | ✓ | | Mime type du fichier | `png` |
| Url | string| `string` |  | | Url du fichier | `/chemin/vers/le/fichier.png` |



### <span id="tag"></span> Tag


> Tag
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` | ✓ | | Nom du Tag | `#python` |



### <span id="user"></span> User


> User
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Address | string| `string` |  | | Address de l'utilisateur
require: false |  |
| Age | uint8 (formatted integer)| `uint8` |  | | Age de l'utilisateur |  |
| Email | string| `string` | ✓ | | Email de l'utilisateur |  |
| FirstName | string| `string` | ✓ | | Nom de l'utilisateur |  |
| LastName | string| `string` | ✓ | | Prénom de l'utilisateur |  |
| Password | string| `string` | ✓ | | Mot de passe de l'utilisateur |  |
| Phone | string| `string` |  | | Numéros de téléphone de l'utilisateur
require: false |  |
| Role | string| `string` | ✓ | | Role de l'utilisateur |  |


