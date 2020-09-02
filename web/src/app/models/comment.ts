// This interface describes a comment as received from the backend
export interface Comment {
  Author: string;
  Created: string;
  Id: string;
  Message: string;
  MessageInsecure: string;
  ThreadId: string;
  ThreadIdInsecure: string;
}
