// Pangaea is a pre-processor for any text file.  Statements and instructions are
// written in JavaScript.  Script blocks are contained within <script>...</script> tags
// and the <%= code %> inline script will cause the output to be written in place.
//
// Example
//
//     <script>
//       function name() {
//         return "Pangaea";
//       }
//     </script>
//     Hello from <%= name() %>.
//
// Produces:
//
//     Hello from Pangaea.
package pangaea
