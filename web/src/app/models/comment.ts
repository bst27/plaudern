// This interface describes a comment as received from the backend
export interface Comment {
  Author: string;
  Created: string;
  Id: string;
  Message: string;
  ThreadId: string;
}
