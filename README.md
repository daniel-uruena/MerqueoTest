# MerqueoTest
___

## Requisitos
- Lenguaje: **GO 1.13.6**
- Servidor de base de datos: **Mongo**

### 1. Configuración de entorno
Despues de clonar el repositorio debes:
- Establecer el directorio `src` como el principal del proyecto
- Instalar las dependencias especificadas en `deps.txt`.
- Restaurar el dump de la base de datos en **MONGO** usando el comando `mongorestore` sobre el directorio `dump-database`

### 2. Iniciar Proyecto
Para iniciar el proyecto ubicarse en el directorio `src` y ejecutar `go run .`.

Si todo funciona bien la aplicación expondrá el API a través de la ruta [http://localhost:5000/merqueo/api](http://localhost:5000/merqueo/api).

Los datos de configuración como los del servidor de base de datos, puerto y prefijo del API se puede configurar en el archivo `appsetting.json` (Se debe recompilar la aplicación para que los cambios sean aplicados).

### 3. EndPoints
Teniendo en cuenta la siguiente URL base, éstos son los endpoints del API:
- UrlBase: http://localhost:5000/merqueo/api.
#### 3.1 Inventario
- **Ruta:** /inventory
- **Método:** GET
- **Descripción:** Permite consultar el inventario completo y lo retorna en forma de lista.
- **Respuesta:** 
```json
[
    {
        "id": 1,
        "date": "2019-03-01",
        "quantity": 3
    },
    {
        "id": 2,
        "date": "2019-03-01",
        "quantity": 3
    },
    ...
]
```

- **Ruta:** /inventory/{idProduct}
- **Método:** GET
- **Descripción:** Permite consultar el inventario de un producto específico (`idProduct`) y lo retorna en forma de lista.
- **Respuesta:** 
```json
[
    {
        "id": 1,
        "date": "2019-03-01",
        "quantity": 3
    },
    {
        "id": 1,
        "date": "2019-03-02",
        "quantity": 1
    },
    ...
]
```

#### 3.2 Proveedores

- **Ruta:** /provider
- **Método:** GET
- **Descripción:** Permite consultar todos los proveedores y lo retorna en forma de lista.
- **Respuesta:** 
```json
[
    {
            "idProvider": 1,
            "name": "Ruby",
            "products": [
                {
                    "id": 1
                },
                {
                    "id": 2
                },
                ...
            ]
    },
    {
            "idProvider": 2,
            "name": "Raul",
            "products": [
                {
                    "id": 28
                },
                {
                    "id": 47
                },
                ...
            ]
    },
    ...
]
```
- **Ruta:** /provider/{idProvider}
- **Método:** GET
- **Descripción:** Permite consultar un proveedor por su identificador.
- **Respuesta:** 
```json
{
   "idProvider": 1,
   "name": "Ruby",
   "products": [
       {
           "id": 1
       },
       {
           "id": 2
       },
       ...
    ]
}
```

- **Ruta:** /provider/product/{idProduct}
- **Método:** GET
- **Descripción:** Permite consultar todos los proveedores que puedan surtir un producto específico (`idProduct`) y lo retorna en forma de lista.
- **Respuesta:** 
```json
[
    {
            "idProvider": 1,
            "name": "Ruby",
            "products": [
                {
                    "id": 1
                },
                {
                    "id": 2
                },
                ...
            ]
    },
    {
            "idProvider": 2,
            "name": "Raul",
            "products": [
                {
                    "id": 28
                },
                {
                    "id": 47
                },
                ...
            ]
    },
    ...
]
```

#### 3.3 Pedidos
- **Ruta:** /order
- **Método:** GET
- **Descripción:** Permite consultar todos los pedidos
- **Respuesta:** 
```json
[
    {
        "id": 1,
        "priority": 1,
        "address": "KR 14 # 87 - 20 ",
        "user": "Sofia",
        "products": [
            {
                "id": 1,
                "name": "Leche",
                "quantity": 1
            },
            {
                "id": 2,
                "name": "Huevos",
                "quantity": 21
            },
            ...
        ]
    },
    {
        "id": 3,
        "priority": 3,
        "address": "KR 13 # 74 - 38 ",
        "user": "Hocks",
        "products": [
            {
                "id": 7,
                "name": "Cebolla Cabezona Blanca Limpia",
                "quantity": 4
            },
            {
                "id": 8,
                "name": "Habichuela",
                "quantity": 3
            },
            ...
        ]
    },
    ...
]
```

- **Ruta:** /order/{idOrder}
- **Método:** GET
- **Descripción:** Permite consultar un pedido
- **Respuesta:** 
```json
{
    "id": 1,
    "priority": 1,
    "address": "KR 14 # 87 - 20 ",
    "user": "Sofia",
    "products": [
        {
            "id": 1,
            "name": "Leche",
            "quantity": 1
        },
        {
            "id": 2,
            "name": "Huevos",
            "quantity": 21
        },
        ...
    ]
}
```

- **Ruta:** /order/product/{idProduct}
- **Método:** GET
- **Descripción:** Permite consultar todos los pedidos que contengan un producto específico (`idProduct`) y los retorna en forma de lista.
- **Respuesta:** 
```json
[
    {
        "id": 1,
        "priority": 1,
        "address": "KR 14 # 87 - 20 ",
        "user": "Sofia",
        "products": [
            {
                "id": 1,
                "name": "Leche",
                "quantity": 1
            },
            {
                "id": 2,
                "name": "Huevos",
                "quantity": 21
            },
            ...
        ]
    },
    {
        "id": 3,
        "priority": 3,
        "address": "KR 13 # 74 - 38 ",
        "user": "Hocks",
        "products": [
            {
                "id": 7,
                "name": "Cebolla Cabezona Blanca Limpia",
                "quantity": 4
            },
            {
                "id": 8,
                "name": "Habichuela",
                "quantity": 3
            },
            ...
        ]
    },
    ...
]
```

#### 3.4 Transportadores
- **Ruta:** /transporter
- **Método:** GET
- **Descripción:** Permite consultar los transportadores requeridos para entregar los pedido y qué pedidos debe entregar cada uno.
- **Respuesta:** 
```json
[
    {
        "idTransporter": 1,
        "orders": [
            {
                "id": 1,
                "priority": 1,
                "address": "KR 14 # 87 - 20 ",
                "user": "Sofia",
                "products": [
                    {
                        "id": 1,
                        "name": "Leche",
                        "quantity": 1
                    },
                    {
                        "id": 2,
                        "name": "Huevos",
                        "quantity": 21
                    },
                    ...
                ]
            },
            ...
        ]
    },
    ...
]
```

#### 3.5 Estadísticas
- **Ruta:** /statistics/bestsold/{deliveryDate}
- **Método:** GET
- **Descripción:** Permite consultar los productos mas vendidos en una fecha específica y los retorna en una lista ordenada de mayor a menor venta.
- **Respuesta:** 
```json
[
    {
        "id": 16,
        "name": "Acelga",
        "quantity": 1500
    },
    {
        "id": 5,
        "name": "Pimentón Rojo",
        "quantity": 100
    },
    ...
]
```

- **Ruta:** /statistics/lesssold/{deliveryDate}
- **Método:** GET
- **Descripción:** Permite consultar los productos menos vendidos en una fecha específica y los retorna en una lista ordenada de menor a mayor venta.
- **Respuesta:** 
```json
[
    {
        "id": 1,
        "name": "Leche",
        "quantity": 1
    },
    {
        "id": 33,
        "name": "Brócoli",
        "quantity": 1
    },
    ...
]
```

#### 3.6 Verificación de inventario
- **Ruta:** /checkinventory/{idOrder}
- **Método:** GET
- **Descripción:** Verifica el inventario disponible para suplir una orden y retorna los productos que deben ser solicitados al proveedor y la cantidad necesaria.
- **Respuesta:** 
```json
{
    "ReadyToDeliver": [
        {
            "id": 5,
            "name": "Pimentón Rojo",
            "quantity": 10
        },
        {
            "id": 6,
            "name": "Kiwi",
            "quantity": 15
        }
    ],
    "NeedToBeRequested": [
        {
            "idProvider": 1,
            "name": "Ruby",
            "products": [
                {
                    "id": 5,
                    "name": "Pimentón Rojo",
                    "quantity": 90
                }
            ]
        },
        {
            "idProvider": 3,
            "name": "Angelica",
            "products": [
                {
                    "id": 6,
                    "name": "Kiwi",
                    "quantity": 45
                }
            ]
        }
    ],
    "Order": {
        "id": 2,
        "priority": 1,
        "address": "KR 20 # 164A - 5 ",
        "user": "Angel",
        "products": [
            {
                "id": 5,
                "name": "Pimentón Rojo",
                "quantity": 100
            },
            {
                "id": 6,
                "name": "Kiwi",
                "quantity": 60
            }
        ],
        "deliveryDate": "2019-03-01"
    }
}
```

#### 3.7 Calculo de inventario
- **Ruta:** /calculateinventory/{date}
- **Método:** GET
- **Descripción:** Calcula y retorna el inventario disponible para la fecha (`date`) teniendo en cuenta el inventario y los pedidos del día anterior.
- **Respuesta:** 
```json
[
    {
        "id": 26,
        "date": "2019-03-02",
        "quantity": 18
    },
    {
        "id": 32,
        "date": "2019-03-02",
        "quantity": 5
    },
    {
        "id": 36,
        "date": "2019-03-02",
        "quantity": 5
    }
]
```

#### 3.8 Errores
En todo caso los errores serán mostrados con la siguiente estructura:
```json
{
    "message": "Mensaje de error"
}
```

### Estructura de la aplicación
Implementé el diseño basado en MVC donde mi capa de "vista" son los end points del API:

- **context.go:**
Encapsula los datos del contexto de ejecución de la aplicación, sirve como punto de control de inyección de dependencias.
- **routes.go:**
Genera las rutas sobre las que expone los end points del API representa en este caso la "vista" que se expone al usuario.
- **Controllers:**
Capa que encapsula la lógica de negocio de la aplicación usando llamados a la capa de "Repositorio".
- **Repositories:**
Capa que encapsula la comunicación con la base de datos diseñada bajo el patrón de fábrica de métodos para que la interfaz de cada modelo pueda ser implementada por estructuras que se comuniquen con distintas bases de datos.
- **Models:**
Capa que contiene las estructuras representativas de las entidades internas y de negocio.
- **tests:**
Contiene las pruebas unitarias y de integración del proyecto