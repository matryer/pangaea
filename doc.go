// Pangaea is a pre-processor for any text file.  Statements and instructions are
// written in JavaScript.  Script blocks are contained within <pangaea type="text/javascript">...</pangaea> tags
// and the <%= code %> inline script will cause the output to be written in place.
//
// Example
//
//     <pangaea type="text/javascript">
//       function name() {
//         return "Pangaea";
//       }
//     </pangaea>
//     Hello from <%= name() %>.
//
// Produces:
//
//     Hello from Pangaea.
package pangaea
