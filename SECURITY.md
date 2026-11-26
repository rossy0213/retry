# Security Policy

## Supported Versions

Use this section to tell people about which versions of your project are
currently being supported with security updates.

| Version | Supported          |
| ------- | ------------------ |
| 5.1.x   | :white_check_mark: |
| 5.0.x   | :x:                |
| 4.0.x   | :white_check_mark: |
| < 4.0   | :x:                |

## Reporting a Vulnerability

Use this section to tell people how to report a vulnerability.

Tell them where to go, how often they can expect to get an update on a
reported vulnerability, what to expect if the vulnerability is accepted or
declined, etc.
# Tarjeta de aplicación: agentes de GitHub Copilot

Obtenga información sobre cómo usar GitHub Copilot características agenticas de forma responsable mediante la comprensión de sus propósitos, funcionalidades y limitaciones.

## ¿Qué es una tarjeta de aplicación?

<!-- Do not modify this text.-->

GitHub’s application and platform cards are intended to help you understand how our AI technology works, the choices application owners can make that influence application performance and behavior, and the importance of considering the whole application, including the technology, the people, and the environment. Application cards are created for AI applications and platform cards are created for AI platform services. These resources can support the development or deployment of your own applications and can be shared with users or stakeholders impacted by them.

As part of its commitment to responsible AI, GitHub adheres to Microsoft's [six core principles](https://www.microsoft.com/en-us/ai/principles-and-approach/?msockid=3da790040c776d6f2b5485e40de56c06#ai-principles): fairness, reliability and safety, privacy and security, inclusiveness, transparency, and accountability. These principles are embedded in the [Responsible AI Standard](https://cdn-dynmedia-1.microsoft.com/is/content/microsoftcorp/microsoft/final/en-us/microsoft-brand/documents/Microsoft-Responsible-AI-Standard-General-Requirements.pdf?culture=en-us\&country=us), which guides teams in designing, building, and testing AI applications. Application and Platform Cards play a key role in operationalizing these principles by offering transparency around capabilities, intended uses, and limitations. For further insight, readers are encouraged to explore Microsoft’s [Responsible AI Transparency Report](https://cdn-dynmedia-1.microsoft.com/is/content/microsoftcorp/microsoft/msc/documents/presentations/CSR/Responsible-AI-Transparency-Report-2025-vertical.pdf) and [Términos de GitHub](/es/site-policy/github-terms).

## 1. Información general

GitHub Copilot incluye varias características agente que van más allá de la sugerencia y la conversación, pueden revisar el código, tomar medidas en su nombre y compilar aplicaciones. Esta tarjeta cubre las siguientes experiencias:

* **Revisión de código de Copilot**: revisa los cambios y metadatos de las solicitudes de extracción en GitHub.com, y genera comentarios de revisión y cambios sugeridos.
* **Agente en la nube de Copilot**: Un agente asíncrono en GitHub.com que puede crear ramas, escribir código y abrir solicitudes de incorporación de cambios en respuesta a incidencias asignadas. El agente en la nube se ejecuta en un entorno efímero y con firewall con análisis de seguridad automatizado.
* **Copilot Cloud**: Copilot en la nube (GitHub Copilot en GitHub.com, GitHub Mobile o en cualquier lugar en el que se seleccione "Nube" como entorno de sesión) puede responder a preguntas sobre los repositorios, planear y realizar cambios, revisar o crear solicitudes de incorporación de cambios y automatizar tareas de repositorio. Una única interfaz basada en el SDK de Copilot combina el chat y el trabajo agéntico.
* **Copilot CLI**: una herramienta de línea de comandos que puede crear y modificar archivos, ejecutar comandos y realizar tareas de varios pasos. De forma predeterminada, el acceso al sistema de archivos se limita al directorio donde se inició la CLI. Los avisos de permisos dependen del modo de permiso activo. Cuando el uso del equipo está habilitado, Copilot CLI también puede interactuar con aplicaciones de escritorio fuera de ese ámbito de directorio.
* **Copilot SDK**: una biblioteca de programación que permite a los desarrolladores crear aplicaciones personalizadas con tecnología de inteligencia artificial mediante Copilot. El SDK se comunica con Copilot CLI a través de JSON-RPC y admite agentes personalizados, integraciones de servidor MCP, enlaces de ciclo de vida y administración de sesiones.
* **La GitHub Copilot app:** una aplicación de escritorio para dirigir las sesiones del agente entre repositorios locales, worktrees de Git y espacios aislados en la nube. La aplicación usa GitHub Copilot CLI y GitHub Copilot SDK como base y agrega una interfaz de usuario para sesiones paralelas, chats rápidos, flujos de trabajo de solicitudes de incorporación de cambios y salida de GitHub, automatizaciones y lienzos. Cuando se habilita el uso del equipo, la aplicación también puede interactuar con las aplicaciones de escritorio.

Estas características comparten principios comunes( supervisión humana, revisión de salidas y uso responsable), pero difieren en sus entornos de ejecución, permisos y flujos de datos. En las secciones siguientes se describe cada una de ellas en su contexto.

## 2. Términos clave

En la lista siguiente se proporciona un glosario de términos clave relacionados con GitHub Copilot Agentes:

* **Sugerencia de código**: cambio de código específico propuesto por la revisión de código de Copilot como parte de la retroalimentación sobre una solicitud de incorporación de cambios. Las sugerencias de código se presentan como cambios sugeridos que se pueden aplicar con un par de clics.
* **Filtrado de contenido**: un sistema de seguridad que examina las solicitudes y respuestas para detectar y bloquear contenido dañino, ofensivo o no seguro antes de que se muestre al usuario.
* **Uso del equipo**: una funcionalidad opcional en GitHub Copilot CLI y el GitHub Copilot app que permite que un agente interactúe con las aplicaciones de escritorio. Sus herramientas pueden leer el contenido de la aplicación accesible y el contexto visual a través del árbol de accesibilidad o las capturas de pantalla del sistema operativo cuando se necesita contexto visual, hacer clic en controles, entrar y editar texto, presionar teclas, desplazarse, arrastrar y navegar por flujos de trabajo entre aplicaciones.
* **Instrucciones personalizadas**: descripciones en lenguaje natural del estilo de codificación y las prácticas recomendadas que un mantenedor de un repositorio puede configurar para orientar los comentarios de la revisión de código de Copilot. Las instrucciones personalizadas ayudan a Copilot comprender las convenciones y estándares de un código base específico.
* **Hallucinación**: un fenómeno en el que un modelo de lenguaje genera una salida que suena factible, pero es incorrecta, no admitida por el contexto proporcionado o completamente fabricada. En la revisión de código, la alucinación puede manifestarse como comentarios que resaltan problemas que no existen o se basan en malentendidos del código.
* **Modelo de lenguaje grande (LLM):** un tipo de red neuronal entrenada en un gran cuerpo de datos de texto que puede generar, analizar y transformar el lenguaje natural y el código. Los agentes de Copilot utilizan uno o varios LLM para procesar el contexto y generar respuestas.
* **Entorno de desarrollo efímero**: un entorno de proceso temporal y aislado creado para cada sesión del agente en la nube. El entorno se destruye una vez finalizada la sesión, lo que garantiza que ningún estado persiste entre ejecuciones.
* **Firewall**: un control de nivel de red habilitado de forma predeterminada para el agente en la nube que impide las conexiones salientes a hosts no autorizados, protegiendo contra la filtración accidental o malintencionada de código o datos.
* **crédito de IA**: unidad de consumo para las funciones de Copilot. Cada uso de la revisión de código de Copilot consume créditos de IA.
* **Permission prompt**: un paso de confirmación interactivo en Copilot CLI que pide al usuario que apruebe una acción(por ejemplo, modificar un archivo, ejecutar un comando o acceder a archivos fuera del directorio actual) antes de que el agente continúe. Las solicitudes de permisos son un mecanismo de seguridad clave para la ejecución de agentes local.
* **Copilot SDK**: una biblioteca de programación (`@github/copilot-sdk`) que permite a los desarrolladores crear aplicaciones personalizadas con tecnología de inteligencia artificial mediante la creación de sesiones, el envío de mensajes y la recepción de respuestas de streaming de Copilot. El SDK se comunica con Copilot CLI a través de JSON-RPC.
* **JSON-RPC**: el protocolo de comunicación usado entre el SDK de Copilot y la CLI de Copilot. El SDK envía solicitudes estructuradas al proceso de la CLI, que controla la interacción del modelo y la ejecución de herramientas.
* **Agente personalizado (SDK)**: una configuración de agente con un nombre definido dentro del SDK que tiene su propia instrucción del sistema, herramientas de alcance limitado y servidores MCP opcionales. El entorno de ejecución del SDK puede delegar automáticamente a subagentes en función de la intención del usuario.
* **Hooks (SDK)**: funciones de devolución de llamada del ciclo de vida en el SDK de Copilot que permiten a los desarrolladores incorporar lógica personalizada en momentos concretos de una sesión, como antes o después de usar herramientas, al iniciar o finalizar la sesión y cuando se produce un error.

## 3. Características o funcionalidades clave

Las características y capacidades clave que se describen aquí explican para qué están diseñados los agentes de GitHub Copilot y cómo funcionan en las tareas compatibles.

* \*\* Comentarios de revisión de códigoautomatizado\*\*: Cuando un usuario solicita una revisión de Copilot, examina los cambios de código más contexto pertinente adicional y proporciona comentarios sobre el código. Los comentarios pueden incluir comentarios de lenguaje natural y sugerencias de código específicas vinculadas a líneas y archivos concretos.
* \*\* Guía de revisión personalizada\*\*: Copilot revisión de código se puede personalizar con instrucciones personalizadas(descripciones de lenguaje natural del estilo de codificación y procedimientos recomendados) para que los comentarios reflejen las convenciones y estándares de un repositorio.
* **Creación de pull requests**: Copilot puede recoger una tarea de una incidencia o un comentario de pull request, crear una rama, generar cambios de código personalizados y abrir una pull request. Una vez creada la solicitud de incorporación de cambios inicial, el agente puede iterar basándose en sus comentarios y revisiones.
* **Ejecución efímera y con firewall**: al trabajar en una tarea, una sesión en la nube puede acceder a su propio entorno de desarrollo efímero, donde puede realizar cambios en el código, ejecutar pruebas automatizadas y ejecutar linters. Un firewall está habilitado de forma predeterminada para evitar la filtración de datos.
* **Examen de seguridad automatizado**: durante la generación de código, el agente en la nube analiza automáticamente el código recién generado para detectar vulnerabilidades de seguridad mediante CodeQL, análisis de secretos y análisis de dependencias, e intenta resolver los problemas antes de que se introduzcan.
* \*\* Integracionesexternal\*\*: El agente en la nube puede recibir información y contexto de MCP como workIQ y Microsoft 365, y aplicaciones externas como Microsoft Teams, Linear, Slack y Jira, lo que permite a los equipos asignar tareas y realizar un seguimiento del progreso directamente dentro de sus flujos de trabajo existentes.
* \*\* Ejecución del agente local (Copilot CLI)\*\*: Copilot CLI proporciona una interfaz similar a chat en el terminal que puede crear y modificar archivos de forma autónoma, ejecutar comandos y realizar tareas de varios pasos. De forma predeterminada, el acceso al sistema de archivos se limita al directorio donde se inició la CLI. El modo de permisos activo determina si la CLI solicita confirmación antes de modificar archivos o ejecutar comandos.
* **Interacción con aplicaciones de escritorio (uso del equipo)**: En las sesiones locales de Copilot CLI y Copilot app, el uso del equipo puede leer el contenido accesible de la aplicación y el contexto visual a través del árbol de accesibilidad del sistema operativo o de capturas de pantalla cuando se necesita contexto visual, hacer clic en controles, introducir y editar texto, pulsar teclas, desplazarse, arrastrar y navegar por flujos de trabajo entre aplicaciones.

## 4. Usos previstos

GitHub Copilot se puede usar en varios escenarios en una variedad de sectores. Algunos ejemplos de casos de uso incluyen:

* **Como complemento de la revisión de código humana**: la revisión de código de Copilot está pensada para ofrecer rápidamente comentarios sobre el código de un desarrollador, lo que permite a los desarrolladores dejar el código listo para su fusión más rápidamente y aumentar la calidad general del código.
* **Mantenimiento de código base**: el agente en la nube puede abordar correcciones relacionadas con la seguridad, actualizaciones de dependencias y refactorización dirigida.
* **Desarrollo de** características: el agente en la nube puede implementar solicitudes de características incrementales, desarrollar conjuntos de pruebas adicionales y crear o actualizar documentación.
* **Creación de prototipos de nuevos proyectos**: El agente en la nube y la CLI de Copilot pueden desarrollar conceptos desde cero, ayudando a los desarrolladores a explorar ideas rápidamente.
* **Setting up your environment (CLI)**: Copilot CLI puede ejecutar comandos en el terminal para configurar el entorno local para que funcione en proyectos existentes.
* **Finding the right command (CLI)**: Copilot CLI puede sugerir comandos para realizar tareas que está intentando completar y explicar comandos desconocidos en lenguaje natural.
* **Trabajar con aplicaciones heredadas y solo con interfaz gráfica (uso del ordenador)**: el uso del ordenador puede ayudar con flujos de trabajo en aplicaciones de escritorio locales que no proporcionan una API, una interfaz de línea de comandos o una integración de MCP. Los usuarios deben describir el resultado, las aplicaciones y las restricciones importantes y, a continuación, revisar las acciones y el resultado.
* **Building custom AI applications (SDK)**: The Copilot SDK permite a los desarrolladores crear aplicaciones que aprovechan Copilot para la generación de código, la interacción del lenguaje natural y la automatización de tareas en sus propios productos y flujos de trabajo.
* **Orquestación multiagente (SDK):** mediante agentes personalizados y subagentes, los desarrolladores pueden crear flujos de trabajo sofisticados en los que varios agentes especializados colaboran en tareas complejas, con delegación automática basada en la intención del usuario.
* **Extensión de aplicaciones con herramientas externas (SDK):** la compatibilidad con el servidor MCP del SDK permite a los desarrolladores conectar sus aplicaciones a servicios y orígenes de datos externos, ampliando la gama de tareas que pueden realizar sus agentes.

## 5. Modelos y datos de entrenamiento

GitHub Copilot aprovecha una variedad de modelos de inteligencia artificial para impulsar la experiencia que ven los usuarios. Para obtener una comparación de los modelos disponibles para Copilot, consulte [Comparación de modelos de IA](/es/copilot/reference/ai-models/model-comparison). Para obtener la lista completa de modelos admitidos, consulte [Modelos de IA admitidos en GitHub Copilot](/es/copilot/reference/ai-models/supported-models). Para obtener información sobre dónde se hospedan los modelos, consulte [Hospedaje de modelos para GitHub Copilot](/es/copilot/reference/ai-models/model-hosting). Para obtener más información sobre los datos utilizados para entrenar los modelos fundacionales en los que se basan GitHub Copilot Agentes, consulte la comparación enlazada de modelos de IA anterior y [¿Con qué datos se ha entrenado GitHub Copilot?](https://github.com/features/copilot#faq) en las Preguntas frecuentes de GitHub Copilot.

La revisión de código de Copilot es un producto diseñado específicamente para este fin que utiliza una combinación cuidadosamente ajustada de modelos, indicaciones y comportamientos del sistema para ofrecer retroalimentación coherente y de alta calidad en una amplia variedad de bases de código. No se admite el cambio de modelo, ya que es probable que el cambio del modelo ponga en peligro la confiabilidad, la experiencia del usuario y la calidad de los comentarios de revisión. La revisión de código de Copilot puede usar modelos que no están habilitados en la página de configuración "Models" de la organización.

Copilot en la nube usa un modelo de lenguaje grande para razonar sobre tareas, generar código y aprovechar las herramientas dentro de su entorno de desarrollo efímero. El agente se ha evaluado en una variedad de lenguajes de programación. Inglés es el idioma principal admitido para las solicitudes y respuestas.

Copilot CLI usa un modelo de lenguaje grande para razonar sobre tareas, generar código, modificar archivos y ejecutar comandos en el entorno de terminal local. El agente se ha evaluado en una variedad de lenguajes de programación. Inglés es el idioma principal admitido para las solicitudes y respuestas.

El SDK de Copilot se comunica con Copilot CLI a través de JSON-RPC, con los mismos modelos y funcionalidades subyacentes. Las aplicaciones compiladas con el SDK usan los mismos modelos disponibles para el usuario o la organización autenticados Copilot. Los desarrolladores también pueden traer sus propias claves de API (BYOK) para usar proveedores de modelos personalizados.

## 6. Rendimiento

### Diferencias por experiencia

#### Revisión del código de Copilot

La revisión de código de Copilot analiza tu código y proporciona comentarios mediante una combinación de procesamiento del lenguaje natural y aprendizaje automático. Este proceso se puede dividir en varios pasos:

1. **Procesamiento de entrada**: los cambios de código se combinan con otra información contextual relevante (por ejemplo, el título y el cuerpo de la solicitud de incorporación de cambios) y las instrucciones personalizadas que se han definido para formar un mensaje. Ese mensaje se envía a un modelo de lenguaje grande.
2. **Análisis del modelo de lenguaje**: El mensaje se procesa mediante el modelo de lenguaje de Copilot, que es una red neuronal que ha sido entrenada con un gran volumen de datos de texto. El modelo de lenguaje analiza la solicitud de entrada.
3. **Generación de respuestas**: el modelo de lenguaje genera una respuesta basada en su análisis del mensaje de entrada. Esta respuesta puede adoptar el formato de sugerencias en lenguaje natural y sugerencias de código.
4. **Formato de salida**: La respuesta se presenta al usuario, ya sea directamente en un editor compatible o como una revisión de una solicitud de extracción en GitHub.com, con comentarios sobre el código vinculados a líneas concretas de archivos concretos. Cuando Copilot ha proporcionado una sugerencia de código, la sugerencia se presenta como un cambio sugerido, que se puede aplicar con un par de clics.

#### CLI de Copilot

Copilot CLI funciona mediante una combinación de procesamiento de lenguaje natural y aprendizaje automático para comprender la tarea y realizar cambios en un código base. Este proceso se puede dividir en varios pasos:

1. **Procesamiento de** entrada: la entrada se combina con información contextual relevante para formar un mensaje. Ese mensaje se envía a un modelo de lenguaje grande para su procesamiento. Las entradas pueden adoptar la forma de lenguaje natural sin formato, fragmentos de código o referencias a archivos del terminal.
2. **Análisis del modelo** de lenguaje: la solicitud se pasa a través de un modelo de lenguaje grande, que es una red neuronal que se ha entrenado en un gran cuerpo de datos. El modelo de lenguaje analiza el prompt de entrada para ayudar al agente a razonar sobre la tarea y a utilizar las herramientas necesarias.
3. **Generación de respuestas**: el modelo de lenguaje genera una respuesta basada en su análisis del mensaje. Esta respuesta puede adoptar la forma de sugerencias de lenguaje natural, sugerencias de código, modificaciones de archivos y ejecuciones de comandos.
4. **Formato de salida**: la respuesta se formatea y se presenta mediante el resaltado de sintaxis, la sangría y otras opciones de formato. El agente también puede ejecutar comandos en el entorno local y crear, editar o eliminar archivos en el sistema de archivos para completar la tarea. Estas acciones están sujetas a los permisos de la sesión y a las restricciones empresariales aplicables.

Puede proporcionar comentarios al agente después de que devuelva una respuesta en la ventana de chat interactiva. El agente volverá a enviar los comentarios al modelo de lenguaje para su posterior análisis y devolverá una respuesta adicional.

#### SDK de Copilot

El SDK de Copilot proporciona una interfaz programática para las capacidades agénticas de Copilot. Las aplicaciones compiladas con el SDK siguen este proceso:

1. **Creación de** sesión: la aplicación crea una sesión con el SDK, especificando el modelo, el símbolo del sistema, las herramientas disponibles, los agentes personalizados, los servidores MCP y los enlaces. El SDK establece una conexión de JSON-RPC a Copilot CLI.
2. **Envío del mensaje**: la entrada del usuario se envía a la sesión. El SDK envía la instrucción (junto con el contexto de la sesión) a la CLI, que la reenvía al modelo de lenguaje.
3. **Ejecución del agente**: El modelo de lenguaje razona sobre la tarea y puede invocar herramientas, delegar en subagentes o conectarse a servidores MCP. Los enlaces de ciclo de vida se activan en cada fase, lo que permite a la aplicación insertar lógica personalizada.
4. **Transmisión en flujo de respuestas**: Las respuestas se envían a la aplicación, que puede presentarlas en cualquier formato adecuado para su interfaz. El SDK proporciona eventos estructurados para texto, llamadas a herramientas, errores y señales de finalización.

## 7. Limitaciones

Comprender las limitaciones de las funciones agénticas de GitHub Copilot es fundamental para garantizar que se utilicen dentro de unos límites seguros y eficaces. Aunque animamos a los clientes a aprovechar estas características en sus soluciones o aplicaciones innovadoras, es importante tener en cuenta que no se diseñaron para todos los escenarios posibles. Animamos a los usuarios a hacer referencia a [Términos de GitHub](/es/site-policy/github-terms) , así como a las siguientes consideraciones al elegir un caso de uso:

* **Problemas de calidad del código no detectados**: es posible que Copilot no identifique todos los problemas presentes en el código, especialmente cuando los cambios son grandes o complejos. Para asegurarse de que todos los problemas pertinentes se identifiquen y corrijan, la revisión de código de Copilot debe complementarse con una cuidadosa revisión humana del código.
* **Falsos positivos**: la revisión de código de Copilot conlleva el riesgo de generar alucinaciones; puede señalar problemas en el código revisado que no existen o que se basan en interpretaciones erróneas del código. Los comentarios generados por Copilot revisión de código deben revisarse cuidadosamente y considerarse antes de tomar medidas y realizar cambios.
* **Sugerencias de código inexactas o no seguras**: Como parte de sus comentarios, la revisión de código de Copilot puede proporcionar sugerencias de código específicas. El código generado puede parecer válido, pero no puede ser realmente semántico o sintácticamente correcto, o puede que no resuelva correctamente el problema identificado en el comentario. Además, el código generado por Copilot puede contener vulnerabilidades de seguridad u otros problemas. Siempre debe revisar y probar cuidadosamente el código generado por Copilot.
* **Posibles sesgos**: los datos de entrenamiento de Copilot proceden de repositorios de código existentes, que pueden contener sesgos y errores que la herramienta puede perpetuar. Además, Copilot revisión de código puede estar sesgada hacia determinados lenguajes de programación o estilos de codificación, lo que puede dar lugar a comentarios poco óptimos o incompletos.
* **Ámbito limitado (agente en la nube):** el modelo de lenguaje usado por el agente en la nube se ha entrenado en un gran cuerpo de código, pero todavía tiene un ámbito limitado y es posible que no pueda controlar determinadas estructuras de código o lenguajes de programación ocultos. Para cada idioma, la calidad de las sugerencias depende del volumen y la diversidad de datos de entrenamiento de ese idioma.
* **Código inexacto (agente en la nube):** el agente en la nube puede generar código que parezca válido, pero que realmente no sea semántica o sintácticamente correcto, o que no refleje con precisión la intención del desarrollador. Debe revisar y probar cuidadosamente el código generado, especialmente cuando se trabaja con aplicaciones críticas o confidenciales.
* **Riesgos de seguridad (agente en la nube):** el agente en la nube genera código y lenguaje natural en función del contexto de un problema o comentario dentro de un repositorio, lo que puede exponer información confidencial o vulnerabilidades si no se usa cuidadosamente. Debería revisar exhaustivamente todos los resultados generados por el agente antes de fusionarlos.
* **Coincidencias de código público (agente en la nube):** el agente en la nube puede generar código que sea una coincidencia o una coincidencia cercana del código disponible públicamente, incluso si la directiva "Sugerencias que coinciden con código público" está establecida en "Bloquear". Si esto sucede, Copilot mostrará coincidencias en los registros de sesión del agente con un vínculo para mostrar los detalles del código coincidente.
* \*\* Ámbito delimitado (CLI)\*\*: el modelo de lenguaje usado por Copilot CLI se ha entrenado en un gran cuerpo de código, pero todavía tiene un ámbito limitado y puede no ser capaz de controlar determinadas estructuras de código o lenguajes de programación ocultos. Para cada idioma, la calidad de las sugerencias depende del volumen y la diversidad de datos de entrenamiento de ese idioma.
* **Código inexacto (CLI)**: Copilot CLI puede generar código que parezca válido, pero que en realidad no sea semántica ni sintácticamente correcto, o puede que no refleje con precisión la intención del desarrollador. Debe revisar y probar cuidadosamente el código generado, especialmente cuando se trabaja con aplicaciones críticas o confidenciales.
* \*\* Riesgos de seguridad (CLI)\*\*: Copilot CLI genera código y lenguaje natural en función del contexto del entorno local, lo que puede exponer información confidencial o vulnerabilidades si no se usa cuidadosamente. Debe revisar exhaustivamente todos los resultados generados por el agente.
* **Coincidencias con código público (CLI)**: Copilot CLI puede generar código que coincida o casi coincida con código disponible públicamente, incluso si la directiva "Sugerencias que coinciden con código público" está configurada como "Bloquear".
* \*\* Riesgos de ejecución de comandos (CLI)\*\*: se requiere precaución adicional al preguntar o permitir que Copilot CLI ejecute un comando, especialmente con respecto a la posible destrucción de algunos comandos sugeridos. Es posible que encuentre comandos para la eliminación de archivos o el formato del disco duro, lo que puede causar problemas si se usan incorrectamente. En última instancia, es responsable de los comandos ejecutados por Copilot CLI.
* **Interpretación de la interfaz (uso de equipo):** el uso del equipo puede seleccionar el control incorrecto, escribir texto en la ubicación incorrecta o tener dificultades con controles no estándar o dinámicos y flujos de trabajo complejos. Los cambios en el estado de tiempo o ventana pueden producir resultados diferentes, hacer que el equipo use para repetir una acción o evitar que continúe.
* **Acciones no deseadas (uso del equipo)**: las instrucciones ambiguas o el contenido inesperado en pantalla pueden hacer que el uso del equipo realice acciones no deseadas. Estas acciones pueden modificar o exponer datos o afectar a las cuentas conectadas. Los usuarios deben mantener las tareas estrechas, especificar restricciones importantes y supervisar la aplicación activa.
* **Información en pantalla (uso del equipo)**: las ventanas de la aplicación pueden mostrar información confidencial o información sobre otras personas. Los usuarios solo deben habilitar el uso de equipos para aplicaciones y tareas cuyo contenido visible sea adecuado para proporcionar como contexto.
* **Restricciones locales y de plataforma (uso de equipo): el uso del equipo** requiere una sesión local y está disponible en macOS y Windows. En macOS, se requieren permisos de accesibilidad y de grabación de pantalla. Si el complemento, el asistente o el servidor MCP agrupados no están disponibles o desconectados, las acciones de uso del equipo no se pueden ejecutar.
* **Aplicaciones siempre permitidas (uso de equipo)**: una aplicación que un usuario elige permitir siempre se puede controlar en sesiones posteriores sin otra solicitud de nivel de aplicación.
  **Permitir siempre** no es adecuado para todas las aplicaciones, especialmente las aplicaciones que contienen información confidencial o permiten acciones de gran impacto.
* \*\* Limitaciones (SDK)\*\*: Dado que el SDK de Copilot se comunica con Copilot CLI, las aplicaciones compiladas con el SDK heredan las mismas limitaciones del modelo, incluido el ámbito limitado para determinados lenguajes de programación y la posibilidad de generar código inexacto o no seguro.
* **Complejidad del agente personalizado (SDK):** los agentes personalizados, las herramientas o los enlaces configurados incorrectamente pueden producir un comportamiento inesperado. Los desarrolladores son responsables de probar y validar el comportamiento de sus configuraciones de agente personalizadas.
* **Confianza de servidor MCP (SDK):** los servidores MCP conectados a través del SDK pueden exponer herramientas y datos de orígenes externos. Los desarrolladores deben asegurarse de que los servidores MCP conectados sean de confianza, ya que los servidores malintencionados o mal configurados podrían introducir un comportamiento perjudicial o exponer datos confidenciales.
* **Variación del modelo BYOK (SDK)**: Al usar configuraciones de trae tu propia clave con proveedores de modelos de terceros, el comportamiento puede diferir del de los modelos alojados en GitHub. Los desarrolladores son responsables de evaluar la seguridad y la calidad de las respuestas de su proveedor elegido.

## 8. Evaluaciones

<!-- Do not modify this text.-->

Las evaluaciones de rendimiento y seguridad evalúan si las aplicaciones de inteligencia artificial funcionan de forma confiable y segura mediante el examen de factores como la base, la relevancia y la coherencia, a la vez que se identifican los riesgos de generar contenido dañino. Las siguientes evaluaciones se llevaron a cabo con componentes de seguridad ya implementados, que también se describen en [9. Componentes y mitigaciones de seguridad](#9-safety-components-and-mitigations).

### Evaluaciones de rendimiento y calidad

GitHub Copilot se evalúa en sus superficies compatibles mediante una combinación de pruebas comparativas estándar del sector (por ejemplo, SWE-Bench) y conjuntos de evaluación desarrollados internamente. Las tareas comparativas se obtienen de repositorios públicos de código abierto y escenarios sintéticos; no se usan consultas de usuario reales ni código de cliente. Cada evaluación incluye varias ejecuciones independientes para tener en cuenta el no determinismo en las salidas del modelo. Entre las métricas clave se incluyen la tasa de resolución (porcentaje de tareas completadas correctamente), la eficacia del token, la latencia y la confiabilidad de las llamadas a herramientas. Los modelos se vuelven a evaluar cuando se realizan y supervisan las actualizaciones continuamente en producción a través de tasas de error, latencia de respuesta y patrones de uso agregados.

### Métodos de evaluación de rendimiento y calidad

Los nuevos modelos para GitHub Copilot se someten a un proceso de evaluación por etapas antes de la implementación. Los equipos integradores ejecutan conjuntos de pruebas comparativas específicos de su superficie, probando el modelo en tareas de codificación representativas, como correcciones de errores, generación de código y refactorización de varios archivos. Los resultados se revisan con respecto a las líneas base establecidas y los modelos de producción existentes. Los modelos deben cumplir o superar el rendimiento de línea base en métricas clave, como la tasa de resolución, la eficacia del token y la latencia, antes de avanzar a la siguiente fase. Un comité de revisión multifuncional toma una decisión formal de seguir o no seguir antes de que se apruebe cualquier modelo para su despliegue de cara al usuario.

### Evaluaciones de riesgos y seguridad

<!-- Do not modify this text.-->

La evaluación de posibles riesgos asociados con el contenido generado por IA es esencial para protegerse frente a los riesgos de contenido con distintos grados de gravedad. Esto incluye la evaluación de la predisposición de una aplicación de inteligencia artificial hacia la generación de contenido dañino o la prueba de vulnerabilidades a ataques de jailbreak. Para GitHub, realizamos evaluaciones de rendimiento, incluidas las adaptadas para fines de codificación de [Microsoft Foundry](https://learn.microsoft.com/en-us/azure/ai-foundry/concepts/evaluation-evaluators/risk-safety-evaluators):

* Odio e injusticia
* Sexual
* Violencia
* Auto-daño
* Material protegido
* Desbloqueo
* Vulnerabilidad de código

### Datos de evaluación para la calidad y la seguridad

<!-- Do not modify this text.-->

Nuestros datos de evaluación están personalizados para evaluar el rendimiento de las aplicaciones de inteligencia artificial en áreas clave de **seguridad** y **calidad**, simulando escenarios y riesgos reales. Empezamos por identificar los aspectos de evaluación pertinentes basados en la investigación multidisciplinaria y la aportación de expertos. Estas preocupaciones se traducen en objetivos de evaluación específicos y guían la formulación de las métricas de evaluación. Para **safety**, creamos avisos adversarios para evitar respuestas no deseadas o de casos perimetrales, que luego se puntúan mediante anotadores asistidos por IA entrenados para evaluar la alineación con los estándares de GitHub. Para la **calidad**, elaboramos instrucciones basadas en rúbricas adaptadas a escenarios, incluida la evaluación de aplicaciones y agentes de generación aumentada con recuperación (RAG). Los conjuntos de datos se recopilan de diversos orígenes, incluidos los conjuntos de datos sintéticos y públicos, para simular escenarios de usuario reales. Con los conjuntos de datos seleccionados, ambas evaluaciones se someten a refinamiento iterativo y ajuste humano para mejorar la confiabilidad y eficacia de las métricas. Esta metodología constituye la base de evaluaciones repetibles y rigurosas que reflejan cómo los clientes usan las evaluaciones para crear una mejor inteligencia artificial.

### Evaluaciones personalizadas

Las funcionalidades de agente de Copilot se han sometido a red teaming de RAI para identificar y abordar posibles riesgos. Seguimos supervisando la eficacia y la seguridad de estas características a lo largo del tiempo. Para obtener más información, consulte [Microsoft AI Red Team construyendo el futuro de una IA más segura](https://www.microsoft.com/en-us/security/blog/2023/08/07/microsoft-ai-red-team-building-future-of-safer-ai/) en el blog de seguridad de Microsoft.

## 9. Componentes de seguridad y mitigaciones

### Copilot en la nube

* **Controles de escalación de privilegios**: en la nube, Copilot es de solo lectura cuando un usuario tiene acceso de lectura, pero no acceso de escritura, a un repositorio. Los flujos de trabajo de acciones desencadenados por las solicitudes de incorporación de cambios generadas por el agente requieren aprobación de un usuario con acceso de escritura antes de que se ejecuten. El agente filtra los caracteres ocultos que podrían permitir a los usuarios ocultar instrucciones maliciosas en comentarios o en el contenido de las incidencias.
* **Permisos restringidos**: el agente en la nube está restringido a los datos de una sola organización en modo "Chat" y también puede realizar cambios en un repositorio de destino cuando se le solicite. Solo puede insertar a una única rama: la rama de la solicitud de incorporación de cambios existente cuando se desencadena a través de `@copilot`, o, en caso contrario, a una nueva rama `copilot/`. El agente no tiene acceso a los secretos de la organización ni a los secretos del repositorio de Actions; solo se pasan al agente los secretos y las variables añadidos específicamente al entorno `copilot`.
* \*\* Garantizar la rastreabilidad\*\*: las confirmaciones del agente en la nube se crean mediante Copilot, con el humano que inició la tarea marcada como coautor. Esto facilita la identificación del código generado por el agente y quién inició la tarea. Las confirmaciones del agente en la nube están firmadas, por lo que aparecen como "Verificadas" en GitHub. Esto ofrece la seguridad de que los commits fueron realizados por el agente en la nube de GitHub Copilot y no se han modificado. Cada mensaje de confirmación incluye un vínculo a los registros de sesión del agente. Esto le proporciona un vínculo permanente de cualquier confirmación escrita por agente a los registros de sesión completos, por lo que puede comprender por qué Copilot realizó un cambio durante la revisión de código o realizar un seguimiento posterior para fines de auditoría.
* **Firewall para la prevención de filtración de datos**: de forma predeterminada, el agente en la nube tiene habilitado un firewall para evitar la filtración del código u otros datos confidenciales, ya sea accidentalmente o debido a una entrada de usuario malintencionada.
* **Detección automatizada de vulnerabilidades de seguridad**: durante la generación de código, el agente en la nube analiza automáticamente el código recién generado para detectar vulnerabilidades de seguridad e intenta resolverlos. El análisis se realiza mediante CodeQL (para identificar posibles vulnerabilidades y errores), el análisis de secretos (para asegurarse de que los secretos no se introducen) y el análisis de dependencias (para comprobar si hay vulnerabilidades conocidas en dependencias a las que se hace referencia).

### CLI de Copilot

* **Acceso con ámbito de directorio**: Copilot CLI limita el acceso a archivos según los permisos de ruta de acceso de la sesión. El acceso fuera de las ubicaciones permitidas requiere permiso adicional.
* **Permisos para modificaciones de archivos**: Copilot CLI comprueba los permisos antes de modificar los archivos. Cuando se le pida aprobación, asegúrese de que se estén modificando los archivos correctos antes de conceder permiso.
* **Permisos para la ejecución de comandos**: Copilot CLI comprueba los permisos antes de ejecutar comandos. Revise con cuidado los comandos cuando se le pida aprobación.
* **Permisos configurables**: puede conceder Copilot CLI permisos específicos o permisos amplios mediante opciones de línea de comandos como `--allow-tool=[TOOLS...]`, `--allow-all-tools` y `--allow-all` (o su equivalente `/permissions allow-all` de comando con barra para una sesión interactiva). Para obtener más información, vea [Referencia de comandos de la CLI de GitHub Copilot](/es/copilot/reference/copilot-cli-reference/cli-command-reference#command-line-options). Como alternativa a conceder permisos amplios, puede usar assisted permissions introduciendo `/permissions assisted` en una sesión interactiva, estableciendo `"assisted"` en `defaultPermissionMode` en el archivo de configuración de usuario o usando `--assisted-approval` para una ejecución no interactiva. Un modelo evalúa las solicitudes de permisos elegibles y aprueba automáticamente las que considera seguras. Se siguen aplicando los permisos existentes y las reglas de empresa. Las solicitudes que aún necesitan su aprobación se le solicitan en una sesión interactiva estándar y se deniegan en modo Autopilot y otras ejecuciones desatendidas. La revisión del modelo no garantiza que una acción sea segura. Para obtener más información, vea [Permitir que los comandos seguros se ejecuten sin aprobación manual](/es/copilot/concepts/customizations/allowing-safe-commands). El modo Autopilot necesita permisos suficientes para completar la tarea sin pedir aprobación. Para obtener más información, vea [Permitir que GitHub Copilot funcione de forma autónoma](/es/copilot/concepts/agent-capabilities/agent-powered-workflows-and-operations/autopilot).
* \*\* Consideraciones de seguridad\*\*: para obtener más información sobre las prácticas de seguridad al usar Copilot CLI, consulte [CLI de GitHub Copilot](/es/copilot/concepts/copilot-surfaces/copilot-cli#security-considerations).

### Uso de ordenador en GitHub Copilot CLI y el GitHub Copilot app

El uso del ordenador es una funcionalidad local opcional que permite a un agente interactuar con las aplicaciones de escritorio en nombre de un usuario. Al igual que otras características agente, el uso del equipo tiene limitaciones. Dado que interpreta interfaces visuales, los cambios en las interfaces o el contenido inesperado en pantalla pueden provocar resultados incoherentes, interrupciones en el flujo de trabajo o acciones no deseadas. Los siguientes controles permiten a los usuarios restringir y revocar su acceso.

* **Activación desactivada por defecto**: el complemento empaquetado de uso del equipo está deshabilitado de forma predeterminada. Los usuarios deben activarlo antes de que sus herramientas estén disponibles para un agente.
* **Directiva administrada**: los administradores de empresa pueden deshabilitar el uso del equipo a través de la configuración administrada. Los usuarios no pueden anular la directiva empresarial mediante la configuración local. Para obtener más información sobre la configuración, consulte [Configuración administrada por la empresa](/es/copilot/reference/enterprise-administrators/enterprise-managed-settings#featurescomputeruse).
* **Controles de permisos de herramientas**: las solicitudes de la aplicación de uso del equipo usan el sistema de permisos del host. Dependiendo del modo de permiso activo, el anfitrión muestra una solicitud de aprobación o aprueba la solicitud automáticamente. Las reglas de denegación explícita tienen prioridad y se produce un error en una solicitud cerrada si no lo acepta ningún controlador de permisos.
* **Permisos del sistema operativo**: En macOS, la guía de uso del ordenador guía a los usuarios en la concesión del permiso de Accesibilidad para interactuar con los controles y del permiso de Grabación de Pantalla para inspeccionar ventanas cuando se necesita contexto visual.
* **Gestión de aprobaciones de aplicaciones**: en GitHub Copilot app, los usuarios pueden revisar y quitar aprobaciones de aplicaciones guardadas. Al eliminar una aplicación, se elimina la aprobación guardada para sesiones futuras, pero no se revoca el acceso ya concedido en una sesión activa. La configuración de permisos del host determina si las solicitudes posteriores muestran un mensaje o se aprueban automáticamente.
* **Interrupción de sesión**: Los usuarios pueden interrumpir una operación local activa si el equipo comienza a comportarse de forma inesperada. En GitHub Copilot CLI, al presionar <kbd>Esc</kbd> dos veces se interrumpe la operación actual. En la GitHub Copilot app, al hacer clic en **Detener** o presionar <kbd>Esc</kbd> se interrumpe.

### SDK de Copilot

* **Controles de seguridad heredados de la CLI**: El SDK de Copilot se comunica con Copilot CLI y hereda su modelo de permisos y sus controles de seguridad. Las ejecuciones de herramientas y las modificaciones de archivos siguen necesitando permisos adecuados.
* **Supervisión basada en** enlaces: los enlaces de ciclo de vida del SDK (como `onPreToolUse` y `onPostToolUse`) permiten a los desarrolladores implementar comprobaciones de seguridad personalizadas, registro de auditoría y flujos de trabajo de aprobación antes y después de la ejecución de la herramienta.
* **Aislamiento del servidor MCP**: los servidores MCP se ejecutan como procesos independientes. Los desarrolladores pueden controlar qué servidores están disponibles por sesión, lo que limita el ámbito de acceso a herramientas externas.
* **Alcance de la sesión**: Cada sesión del SDK está aislada con su propio contexto, herramientas y permisos. Los desarrolladores pueden controlar qué datos y funcionalidades están disponibles en cada sesión.
* **RESPONSABILIDAD DE BYOK**: al usar configuraciones bring-your-own-key, se envían mensajes y respuestas directamente al proveedor configurado. Los desarrolladores son responsables de revisar las directivas de control de datos de su proveedor elegido.

### El GitHub Copilot app

GitHub Copilot app es una aplicación de escritorio para gestionar las sesiones del agente en repositorios locales, árboles de trabajo de Git y entornos aislados en la nube. La aplicación usa GitHub Copilot CLI y GitHub Copilot SDK como base y agrega una interfaz de usuario para sesiones paralelas, chats rápidos, flujos de trabajo de solicitudes de incorporación de cambios y salida de GitHub, automatizaciones y lienzos.

* **Aislamiento del espacio de trabajo**: las sesiones de la aplicación pueden ejecutarse en árboles de trabajo de Git independientes o en sandboxes en la nube, lo que ayuda a aislar las tareas simultáneas entre sí y de la rama predeterminada.
* **Cli heredada y controles de seguridad del SDK**: las sesiones de la aplicación usan GitHub Copilot CLI y GitHub Copilot SDK, por lo que heredan las mismas limitaciones del modelo y los controles de seguridad para el acceso a archivos, modificación de archivos, ejecución de comandos, ámbito de sesión y herramientas conectadas.
* **Revisión humana antes de combinar**: la aplicación expone los cambios generados, el estado de la solicitud de incorporación de cambios, los resultados de ci, la actividad de revisión y el historial de sesiones para que los usuarios puedan inspeccionar el trabajo antes de enviar revisiones, habilitar Agent Mergeo combinar solicitudes de incorporación de cambios.
* **Integraciones y automatizaciones configurables**: los usuarios controlan qué servidores MCP, aptitudes, extensiones de lienzo y automatizaciones están disponibles en la aplicación. Estas configuraciones deben revisarse antes de su uso, ya que pueden afectar a las herramientas, los datos y las acciones en segundo plano disponibles para las sesiones.

La integración de Azure DevOps en GitHub Copilot app tiene estos controles de acceso y control de datos adicionales:

* **Tratamiento de datos**: El código privado, el contenido de la solicitud de incorporación, los comentarios de revisión y los extractos del registro de compilación devueltos al agente pueden entrar en el contexto del modelo seleccionado. El contenido que aparece en los resultados del chat o de la herramienta se conserva en el historial de sesiones local. Se aplican las restricciones del modelo, la configuración de entrenamiento y las directivas empresariales existentes.
  los modelos proporcionados por GitHub siguen los términos y políticas de Copilot para el plan del usuario. Los usuarios de los modelos de otros proveedores deben revisar los términos y políticas de tratamiento de datos de esos proveedores.
* **Controles de acceso y acción**: La integración está deshabilitada de forma predeterminada y los usuarios deben habilitarla antes de utilizarla. Los administradores pueden deshabilitar la integración, pero no pueden restringirla a repositorios seleccionados a través de su configuración. Los permisos de Azure DevOps se aplican a todas las acciones. Los permisos de la herramienta de sesión (como preguntar cada vez o permitir todos) rigen las acciones del agente, incluida la creación de solicitudes de incorporación de cambios y la respuesta a o la resolución de subprocesos de revisión. Dependiendo de estos permisos, esas acciones pueden continuar sin otra solicitud de aprobación. La vista de solicitud de incorporación de cambios también proporciona controles que el usuario puede activar para aprobaciones, cambios de revisor, comentarios, ediciones, combinaciones y acciones de compilación.
  Agent Merge está deshabilitado por separado de forma predeterminada y debe estar habilitado para cada solicitud de incorporación de cambios. Una vez habilitado, puede completar esa solicitud de incorporación de cambios sin una confirmación independiente, sujeta a las comprobaciones y directivas de Azure DevOps.
* **Revocación de la cuenta**: al cerrar sesión del cuenta Microsoft conectado, se quitan los tokens de autenticación almacenados localmente, pero no se revocan con Microsoft. La revocación se administra a través de controles de cuenta de Microsoft Entra.

### Control de datos al usar su propio proveedor de modelos (CLI)

Al configurar Copilot CLI para usar su propio proveedor de modelos, las indicaciones, el contexto de código y las respuestas generadas se envían directamente al proveedor que configure. No se enrutan a través de GitHub. Usted es responsable de revisar y cumplir con los términos de servicio y las directivas de control de datos de su proveedor elegido.

#### Telemetry

Cuando usa su propio proveedor de modelos sin modo sin conexión, la CLI de Copilot sigue enviando telemetría a GitHub de la forma habitual. Esta telemetría no incluye tus indicaciones ni tu código, pero sí incluye metadatos de uso.

Si habilita el modo sin conexión estableciendo la variable de entorno `COPILOT_OFFLINE` en `true`, se deshabilita toda la telemetría. En modo sin conexión, Copilot CLI solo realiza solicitudes de red al proveedor de modelos configurado.

#### Autenticación y disponibilidad de características

No se requiere la autenticación de GitHub al usar tu propio proveedor de modelos (BYOK). Sin GitHub autenticación, las siguientes características no están disponibles:

* `/delegate`, que transfiere la sesión al Copilot del servidor de GitHub
* Servidor MCP de GitHub
* búsqueda de código de GitHub

En el modo sin conexión, también se deshabilitan las herramientas basadas en web, como `web_fetch` y GitHub Búsqueda de código.

#### Sin recurso a modelos alojados en GitHub

Si la configuración del proveedor de modelos no es válida, Copilot CLI se cierra con un error. No recurre a modelos alojados en GitHub como alternativa. Los errores comunes, como conexión rechazada, errores de autenticación, modelo no encontrado y tiempos de espera agotados, muestran mensajes claros con indicaciones concretas para resolverlos.

## 10. Procedimientos recomendados para implementar y adoptar GitHub Copilot

La inteligencia artificial responsable es un compromiso compartido entre GitHub y sus clientes. Aunque GitHub crea aplicaciones de inteligencia artificial con seguridad, imparcialidad y transparencia en el núcleo, los clientes desempeñan un papel fundamental en la implementación y el uso de estas tecnologías de forma responsable dentro de sus propios contextos. Para admitir esta asociación, ofrecemos los siguientes procedimientos recomendados para los implementadores y los usuarios finales con el fin de ayudar a los clientes a implementar la inteligencia artificial responsable de forma eficaz.

* **Actúe con precaución y evalúe los resultados al usar las funciones autónomas de Copilot para decisiones importantes o en ámbitos sensibles**: <!-- Do not modify this text.-->
  Las decisiones consecuentes son aquellas que pueden tener un impacto legal o significativo en el acceso de una persona a la educación, el empleo, las plataformas financieras, los beneficios gubernamentales, la salud, la vivienda, el seguro, las plataformas legales o que podrían dar lugar a daños físicos, psicológicos o financieros. Los dominios confidenciales, como las plataformas financieras, la atención sanitaria y la vivienda, requieren atención especial debido al potencial de impacto desproporcionada en diferentes grupos de personas. Al usar la inteligencia artificial para tomar decisiones en estas áreas, asegúrese de que las partes interesadas afectadas puedan comprender cómo se toman las decisiones, apelar decisiones y actualizar los datos de entrada pertinentes.
* **Evaluar consideraciones legales y normativas**: <!-- Do not modify this text.-->
  Los clientes deben evaluar posibles obligaciones legales y normativas específicas al usar cualquier plataforma y soluciones de inteligencia artificial, que puede que no sean adecuadas para su uso en todos los sectores o escenarios. Además, las plataformas o soluciones de inteligencia artificial no están diseñadas para y pueden no usarse de maneras prohibidas en términos de servicio aplicables y códigos de conducta pertinentes.
* **Ejercicio de la supervisión humana cuando proceda**: la supervisión humana es una protección importante al interactuar con las aplicaciones de inteligencia artificial. Aunque mejoramos continuamente nuestras aplicaciones de inteligencia artificial, la inteligencia artificial podría cometer errores. Los resultados generados pueden ser inexactos, incompletos, sesgados, desalineados o irrelevantes para los objetivos previstos. Esto podría ocurrir debido a varias razones, como la ambigüedad en las entradas o limitaciones de los modelos subyacentes. Por ello, los usuarios deben revisar las respuestas generadas por las funcionalidades agénticas de Copilot y verificar que cumplen con sus expectativas y requisitos.
* **Tenga en cuenta el riesgo de sobrecarga**: <!-- Do not modify this text.-->
  La dependencia excesiva de la inteligencia artificial se produce cuando los usuarios aceptan salidas de IA incorrectas o incompletas, principalmente porque los errores en las salidas de IA pueden ser difíciles de detectar. Para el usuario final, el exceso de dependencia podría dar lugar a una disminución de la productividad, la pérdida de confianza, el abandono de aplicaciones, la pérdida financiera, daños psicológicos, daños físicos, entre otros. (por ejemplo, un médico acepta una salida de IA incorrecta).
* **Tenga cuidado al diseñar inteligencia artificial agente en dominios confidenciales**: <!-- Do not modify this text.-->
  Los usuarios deben tener precaución al diseñar o implementar aplicaciones de inteligencia artificial agente en dominios confidenciales en los que las acciones del agente son irreversibles o muy consecuentes. También se deben tomar precauciones adicionales al crear inteligencia artificial agente autónoma, tal como se describe más adelante en [autotitle](/es/site-policy/github-terms).
* **Usa la revisión de código de Copilot para complementar las revisiones humanas, no para sustituirlas**: Aunque la revisión de código de Copilot puede ser una herramienta potente para mejorar la calidad del código, es importante utilizarla como complemento, en lugar de sustituir las revisiones humanas. Siempre debes revisar y verificar los comentarios generados por la revisión de código de Copilot, y complementar los comentarios de Copilot con una cuidadosa revisión humana para asegurarte de que tu código cumpla con tus requisitos.
* **Enviar comentarios**: Si encuentra algún problema o alguna limitación con la revisión de código de Copilot, le recomendamos que envíe sus comentarios mediante los botones de pulgar hacia arriba y pulgar hacia abajo en los comentarios de Copilot. Esto puede ayudar a GitHub mejorar la herramienta y abordar cualquier problema o limitaciones.
* **Configurar instrucciones personalizadas**: puede configurar instrucciones personalizadas para ayudar a Copilot comprender el estilo de codificación y los procedimientos recomendados, mejorando la relevancia y la calidad de los comentarios de revisión.
* **Asegúrese de que las tareas estén bien delimitadas**: cuanto más clara y bien delimitada esté la indicación que asigne al agente en la nube, mejor será el resultado. Un problema ideal incluye una descripción clara del problema, criterios de aceptación completos y sugerencias sobre qué archivos deben cambiarse.
* **Personalice el agente en la nube con contexto adicional**: el agente en la nube tiene acceso a la búsqueda de código semántico, lo que le ayuda a encontrar código relevante en función del significado en lugar de simplemente coincidencias exactas de texto, lo que le permite completar tareas más rápido. Para mejorar aún más el rendimiento, implemente instrucciones de Copilot personalizadas para ayudar al agente a comprender mejor el proyecto y a compilar, probar y validar sus cambios.
* **Use el agente en la nube como herramienta, no un reemplazo**: siempre debe revisar y probar el contenido generado por el agente en la nube para asegurarse de que cumple sus requisitos y está libre de errores o problemas de seguridad antes de la combinación.
* **Uso de procedimientos de codificación y revisión de código seguros con el agente en la nube**: aunque el agente en la nube puede generar código sintácticomente correcto, es posible que no siempre sea seguro. Siga los procedimientos recomendados para la codificación segura y la revisión de código. Tome las mismas precauciones que haría con cualquier código que use material que no se originó de forma independiente, incluidas las pruebas rigurosas, el examen de IP y la comprobación de vulnerabilidades de seguridad.
* **Manténgase al día**: el agente en la nube es una tecnología en evolución. Manténgase al día con los nuevos riesgos de seguridad o los procedimientos recomendados que puedan surgir.
* **Use Copilot CLI como herramienta, no un reemplazo**: siempre debe revisar y comprobar los comandos y el código generados por Copilot CLI para asegurarse de que cumplen sus requisitos y están libres de errores o problemas de seguridad.
* **Revise los comandos antes de ejecutarlos (CLI)**: Tenga especial precaución cuando Copilot CLI sugiera ejecutar comandos, especialmente aquellos que modifican o eliminan archivos. En última instancia, es responsable de los comandos que permite ejecutar al agente.
* **Mantén las tareas de la CLI bien delimitadas**: Cuanto más clara y mejor delimitada sea la instrucción que proporciones, mejores serán los resultados. Incluya una descripción clara del problema, los criterios de aceptación y las sugerencias sobre qué archivos deben cambiarse.
* **Enviar comentarios (CLI)**: Si encuentra algún problema o alguna limitación con la CLI de Copilot, envíe sus comentarios con el comando `/feedback`.
* **Usar herramientas directas cuando estén disponibles (uso del ordenador):** cuando una API, un servidor MCP, un comando terminal, una herramienta de sistema de archivos o una herramienta de explorador dedicada pueda completar una tarea, use esa herramienta para obtener información más estructurada y resultados predecibles.
* **Mantenga las tareas de uso del ordenador limitadas y observables**: describa el resultado que desee, las aplicaciones implicadas y las restricciones importantes. Revise la aplicación activa y su resultado.
* **Limitar las aplicaciones siempre permitidas**: evite elegir **Permitir siempre** en aplicaciones que contienen información confidencial o admiten acciones de alto impacto. Revise y quite las aprobaciones guardadas cuando ya no sean necesarias.
* **Validar el comportamiento del agente personalizado (SDK):** pruebe exhaustivamente agentes personalizados, herramientas y enlaces antes de implementar aplicaciones compiladas con el SDK en producción. Asegúrese de que las configuraciones de herramientas y las indicaciones del sistema producen un comportamiento seguro y esperado.
* **Auditar conexiones de servidor MCP (SDK):** solo conéctese a servidores MCP que confíe. Revise las herramientas y los datos que expone cada servidor y asegúrese de que se alinean con los requisitos de seguridad de la aplicación.
* **Implementación de enlaces de seguridad (SDK):** use los enlaces de ciclo de vida del SDK para implementar límites de protección, como el filtrado de contenido, el registro de auditoría y los flujos de trabajo de aprobación de herramientas en las aplicaciones.
* **Defina adecuadamente el alcance de las sesiones (SDK)**: Configure cada sesión del SDK solo con las herramientas, los agentes y los permisos necesarios para la tarea en cuestión. Evite conceder acceso amplio cuando sea suficiente un ámbito estrecho.
* **Revise las políticas del proveedor de BYOK (SDK)**: si utiliza configuraciones de clave propia (BYOK), asegúrese de que las condiciones del servicio y las políticas de tratamiento de datos del proveedor del modelo elegido cumplan los requisitos de su organización.

## 11. Más información sobre las características de agente de GitHub Copilot

Para obtener instrucciones adicionales sobre el uso responsable de las características de Copilot agente, se recomienda revisar la siguiente documentación:

* [Uso de GitHub Copilot para la revisión de código](/es/copilot/how-tos/use-copilot-agents/request-a-code-review/use-code-review)
* [Procedimientos recomendados para usar GitHub Copilot para trabajar en tareas](/es/copilot/tutorials/cloud-agent/get-the-best-results)
* [Configurar el entorno de desarrollo para agentes en Acciones de GitHub](/es/copilot/how-tos/copilot-on-github/customize-copilot/customize-cloud-agent/customize-the-agent-environment)
* [Personalización o deshabilitación del firewall para GitHub Copilot](/es/copilot/how-tos/copilot-on-github/customize-copilot/customize-the-firewall)
* [Crea tu primera aplicación con tecnología Copilot](/es/copilot/get-started/sdk-quickstart)
* [Configuración de servidores MCP para el repositorio](/es/copilot/how-tos/copilot-on-github/customize-copilot/configure-mcp-servers)
* [CLI de GitHub Copilot](/es/copilot/concepts/copilot-surfaces/copilot-cli)
* [aplicación Copilot de GitHub](/es/copilot/concepts/copilot-surfaces/github-copilot-app)
* [Acerca del uso del ordenador en GitHub Copilot](/es/copilot/concepts/agent-capabilities/agent-powered-workflows-and-operations/computer-use)
* [Términos de GitHub para productos y funciones adicionales](/es/site-policy/github-terms/github-terms-for-additional-products-and-features#github-copilot)
* [Centro de confianza de Copilot](https://copilot.github.trust.page/)

### Obtener más información sobre la IA responsable

* [Principios de inteligencia artificial de Microsoft](https://www.microsoft.com/en-us/ai/responsible-ai)
* [Recursos de inteligencia artificial responsable de Microsoft](https://www.microsoft.com/en-us/ai/responsible-ai-resources)
* [Cursos de Microsoft Azure Learning sobre IA responsable](https://docs.microsoft.com/en-us/learn/paths/responsible-ai-business-principles/)
